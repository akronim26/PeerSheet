package main

import (
	"fmt"
	"os"

	"github.com/akronim26/peer-sheet/p2p"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Notice: .env file not found or could not be loaded: %v. Continuing...\n", err)
	}

	err = p2p.RunRelayNode()
	if err != nil {
		fmt.Printf("Error running relay node: %v\n", err)
		os.Exit(1)
	}
}