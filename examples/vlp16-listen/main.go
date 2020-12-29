package main

import (
	"context"
	"fmt"
	"os"

	"github.com/einride/vlp16-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: vlp16-listen <port>")
		os.Exit(1)
	}
	// Bind UDP packet listener.
	listener, err := vlp16.ListenUDP(context.Background(), fmt.Sprintf("0.0.0.0:%s", os.Args[1]))
	if err != nil {
		panic(err)
	}
	// Listen for and print packets.
	for {
		if err := listener.ReadPacket(); err != nil {
			panic(err)
		}
		fmt.Println(listener.SourceIP(), listener.Packet().ProductID)
		fmt.Println(listener.RawPacket())
		fmt.Println()
	}
}
