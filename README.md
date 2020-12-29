# VLP-16 Go

Go SDK for reading and parsing data from [Velodyne][velodyne] [VLP-16
(a.k.a. Puck)][puck] sensors.

[velodyne]: https://velodynelidar.com/
[puck]: https://velodynelidar.com/products/puck/

## Documentation

See the [VLP-16 product page][puck] and the [VLP-16 packet
structure][packet-structure] specification.

[packet-structure]: https://velodynelidar.com/wp-content/uploads/2019/09/63-9276-Rev-C-VLP-16-Application-Note-Packet-Structure-Timing-Definition.pdf

## Installing

```bash
$ go get -u go.einride.tech/vlp16
```

## Examples

### Listen for VLP-16 packets

```go
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
```
