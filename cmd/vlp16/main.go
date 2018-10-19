package main

import (
	"log"

	"github.com/einride/recoverer"
)
import "github.com/einride/vlp-16-go/pkg/vlp16"

const (
	port = 2368
)

func main() {
	recoverer.AutoRecover(func(err interface{}) {
		log.Printf("Connecting to %v, %v", port, err)
		err = vlp16.InitVLP16(port)
		if err != nil {
			log.Printf("recovered: %+v", err)
		}
	})
}
