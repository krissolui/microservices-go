package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	log.Println("Starting listener...")

	// try to connect to RabbitMQ
	rabbitConn, err := connectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConn.Close()

	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func connectRabbitMQ() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err == nil {
			log.Println("Connected to RabbitMQ!")
			connection = conn
			break
		}

		log.Println("RabbitMQ not yet ready...")
		counts++

		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		// increase wait time exponentially: counts ^ 2
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
	}

	return connection, nil
}
