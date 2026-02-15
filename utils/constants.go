package utils

import (
	"time"
)

const (
	ProtocolNegotiationTimeout = 10 * time.Second
	UpgradeTimeout             = 10 * time.Second
	DialTimeout                = 30 * time.Second
	MinConnections			   = 1
	MaxConnections             = 1000
	MaxIncomingPending         = 100 
	MaxPeerAddrsToDial         = 10  
	MaxReservations            = 128
	ReservationTTL             = 1 * time.Hour
	DefaultDataLimit           = 1 << 30 // 1 GB
	DefaultDurationLimit       = 2 * time.Hour
	HopTimeout                 = 30 * time.Second
)
