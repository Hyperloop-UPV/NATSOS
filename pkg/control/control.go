package control

import (
	"log"
	"net"
	"strconv"
)

// Control server is the auxiliar server that controls Chimera's emulator options

func StartControlServer(port int) {

	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))

	// Configure server
	server := NewServer(addr, handleCommand)

	log.Printf("Control server started in 0.0.0.0:%d", port)

	if err := server.Start(); err != nil {
		log.Fatal(err)
		return
	}

}

func handleCommand(cmd Command) string {

	return "si soy"

}
