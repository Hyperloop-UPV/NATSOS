package main

import (
	"flag"
	"log"

	"github.com/Hyperloop-UPV/NATSOS/pkg/adj"
	"github.com/Hyperloop-UPV/NATSOS/pkg/config"
	"github.com/Hyperloop-UPV/NATSOS/pkg/network"
)

func main() {

	// Get the configuration file path from command line arguments
	configFile := flag.String("config", "config.json", "path to the configuration file")
	flag.Parse()

	// Load the configuration
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up the network configuration
	if err := network.SetupNetwork(); err != nil {
		log.Fatalf("Failed to setup network: %v", err)
	}

	err = network.SetupExternalInterface(cfg.Network.Interface, "192.168.0.1/16")
	if err != nil {
		log.Fatalf("Failed to setup external interface: %v", err)
	}
	// get the ADJ branch from the configuration and print it
	adj, err := adj.NewADJ(cfg.ADJBranch, cfg.ADJPath)
	if err != nil {
		log.Fatalf("Failed to initialize ADJ: %v at %s", err, cfg.ADJPath)
	}

	// create a dummy interface for each board in the ADJ
	for _, board := range adj.Boards {

		network.SetupDummyInterface(board.Name, board.IP)
	}

	// Print the ADJ info and boards
	log.Printf("ADJ Info: %+v", adj.Info)
}
