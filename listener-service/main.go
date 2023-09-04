package main

import (
	"listener/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	log.Println("Starting listener...")

	// try to connect to RabbitMQ
	rabbitConn, err := connectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConn.Close()

	// start listening for messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Fatal(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
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
