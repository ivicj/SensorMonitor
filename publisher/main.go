package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		panic(err)
	}

	if !nc.IsConnected() {
		fmt.Println("Publisher Not Connected")
	}

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	type message struct {
		Name      string
		Timestamp int64
		Value     float32
	}

	for {

		messages := []message{
			{"Sensor1", time.Now().UnixMilli(), 52.7},
			{"Sensor2", time.Now().UnixMilli(), 22.3},
			{"Sensor3", time.Now().UnixMilli(), 33.8},
		}

		fmt.Println("Sending...")
		c.Publish("sensor", messages)
		time.Sleep(1 * time.Second)

	}
}
