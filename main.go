package main

import (
	"flag"
	"log"

	"github.com/Hyperloop-UPV/NATSOS/pkg/adj"
	"github.com/Hyperloop-UPV/NATSOS/pkg/config"
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

	// get the ADJ branch from the configuration and print it
	adj, err := adj.NewADJ(cfg.ADJBranch, cfg.ADJPath)
	if err != nil {
		log.Fatalf("Failed to initialize ADJ: %v at %s", err, cfg.ADJPath)
	}

	// Print the ADJ info and boards
	log.Printf("ADJ Info: %+v", adj.Info)
}
