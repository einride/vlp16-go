package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"

	vlp16 "github.com/einride/vlp-16-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: vlp-16-cat <port>")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: port})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client := vlp16.NewClient(udpConn)
	for {
		if err := client.Receive(context.Background()); err != nil {
			panic(err)
		}
		fmt.Println(client.SenderAddr())
		fmt.Println(hex.EncodeToString(client.RawPacket()))
		fmt.Println()
	}
}
