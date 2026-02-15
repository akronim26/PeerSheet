package p2p

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akronim26/peer-sheet/utils"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	webrtc "github.com/libp2p/go-libp2p/p2p/transport/webrtc"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
)

func RunRelayNode() error {
	privKey, peerID, err := utils.LoadIdentityFromEnv()
	if err != nil {
		return fmt.Errorf("failed to load identity: %w", err)
	}

	cm, err := connmgr.NewConnManager(
		utils.MinConnections,
		utils.MaxConnections,
		connmgr.WithGracePeriod(time.Minute),
	)
	if err != nil {
		return fmt.Errorf("failed to create connection manager: %w", err)
	}

	relayResources := relay.DefaultResources()
	relayResources.MaxReservations = utils.MaxReservations
	relayResources.ReservationTTL = utils.ReservationTTL
	relayResources.Limit.Data = utils.DefaultDataLimit
	relayResources.Limit.Duration = utils.DefaultDurationLimit

	host, err := libp2p.New(
		libp2p.Identity(privKey),

		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9091",
			"/ip4/0.0.0.0/tcp/9092/ws",
			"/ip4/0.0.0.0/udp/9093/webrtc-direct",
		),

		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(websocket.New),
		libp2p.Transport(webrtc.New),

		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),

		libp2p.Security(noise.ID, noise.New),

		libp2p.ConnectionManager(cm),

		libp2p.EnableNATService(),
		libp2p.EnableRelayService(relay.WithResources(relayResources)),
		libp2p.Ping(true),
	)
	if err != nil {
		return fmt.Errorf("failed to create libp2p host: %w", err)
	}

	fmt.Printf("Libp2p host started with ID: %s\n", peerID)
	fmt.Printf("Listening on addresses:\n")
	for _, addr := range host.Addrs() {
		fmt.Printf(" - %s\n", addr)
	}

	fmt.Println("Relay node running. Press Ctrl+C to shut down.")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("\nShutting down relay node...")

	return host.Close()
}