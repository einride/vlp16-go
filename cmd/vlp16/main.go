package main

import (
	"fmt"
	"log"
	"net"

	"github.com/einride/vlp-16-go/pkg/vlp16"
)

const (
	port     = 2368
	protocol = "udp"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr(protocol, fmt.Sprintf(":%v", port))
	if err != nil {
		log.Panic("Couldn't resolve address")
	}

	conn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		log.Panic("Couldn't connect")
	}

	if err := vlp16.InitVLP16(conn); err != nil {
		log.Fatalf("Error from VLP16: %v", err)
	}

}
