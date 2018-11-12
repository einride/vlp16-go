package main

import (
	"fmt"
	"log"
	"net"

	"github.com/einride/vlp-16-go"
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

	packet := vlp16.Packet{}
	for {
		cloud := vlp16.SphericalPointCloud{}
		if err := packet.Read(conn); err != nil {
			log.Printf("Error reading from connection. %v", err)
		}

		if err := cloud.UnmarshalPacket(&packet); err != nil {
			log.Panic("Error parsing packet")
		}
	}
}
