package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test1234"
	dbname   = "postgres"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		panic(err)
	}

	if !nc.IsConnected() {
		fmt.Println("Subscriber Not Connected")
	}

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	type message struct {
		Name      string
		Timestamp int64
		Value     float32
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c.Subscribe("sensor", func(m *[]message) {
		fmt.Printf("Received a message: %+v\n", m)

		sensorInfo := *m

		rawDataQuery := `
			INSERT INTO public.SensorRaw ("Name", "Timestamp", "Value")
			VALUES ($1, $2, $3),
				   ($4, $5, $6),
				   ($7, $8, $9)`

		_, err = db.Exec(rawDataQuery,
			sensorInfo[0].Name, sensorInfo[0].Timestamp, sensorInfo[0].Value,
			sensorInfo[1].Name, sensorInfo[1].Timestamp, sensorInfo[1].Value,
			sensorInfo[2].Name, sensorInfo[2].Timestamp, sensorInfo[2].Value)

		if err != nil {
			panic(err)
		}

		avg := (sensorInfo[0].Value + sensorInfo[1].Value + sensorInfo[2].Value) / 3

		avgDataQuery := `
			INSERT INTO public.SensorAverage ("Timestamp", "Value")
			VALUES ($1, $2)`

		_, err = db.Exec(avgDataQuery, sensorInfo[0].Timestamp, avg)
	})

	for {
		fmt.Println("Sleeping...")
		time.Sleep(2 * time.Second)
	}

}
