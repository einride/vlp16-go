package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	vlp16 "github.com/einride/vlp16-go"
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
	client, err := vlp16.ListenUDP(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", port),
		vlp16.WithBatchSize(10),
		vlp16.WithBufferSize(2097152),
	)
	if err != nil {
		panic(err)
	}
	for {
		if err := client.Receive(context.Background()); err != nil {
			panic(err)
		}
		fmt.Println(client.SourceIP())
		fmt.Println(hex.EncodeToString(client.RawPacket()))
		fmt.Println()
	}
}
