package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go.einride.tech/vlp16"
)

func runEmulationFromFile(addr string, file string) error {

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("run emulation from file: %w", err)
	}

	var conn net.Conn
	for {
		if conn == nil {
			log.Println("Connecting to: ", addr)
			conn, err = net.DialTimeout("udp4", addr, 10*time.Second)
			if err != nil {
				log.Println("Dialing connection not possible, retrying: %w", err)
				time.Sleep(500 * time.Millisecond)
				continue
			}
			break
		}
	}

	var rawPackets []vlp16.RawPacket
	sc := bufio.NewScanner(f)
	sc.Split(vlp16.ScanPackets)
	var numPackets int
	for sc.Scan() {
		numPackets++
		var rawPacket vlp16.RawPacket
		copy(rawPacket[:], sc.Bytes())
		rawPackets = append(rawPackets, rawPacket)
	}

	i := 0
	for {

		packet := rawPackets[i]

		i++
		if i >= len(rawPackets) {
			log.Println("End of recorded file. Restarting.")
			i = 0
			continue
		}

		// If we get some error then just wait a while for new connection
		if err = conn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if _, err = conn.Write(packet[:]); err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}

	}

	return nil
}

func main() {
	playFile := flag.String("play", "recording.bin", "play udp packets from file")
	port := flag.Int64("p", 2368, "port to listen on")
	simulation := flag.Bool("s", false, "run vlp16 with pre-recorded data")
	flag.Parse()

	addrPort := fmt.Sprintf(":%d", *port)
	if *simulation {
		log.Printf("Running emulation from file: %s", *playFile)
		if err := runEmulationFromFile(addrPort, *playFile); err != nil {
			log.Panic(err)
		}
		return
	}

	// Bind UDP packet listener.
	listener, err := vlp16.ListenUDP(context.Background(), addrPort)
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
