package main

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	ch.Qos(1, 0, false)

	msgs, err := ch.Consume(
		"order_queue",
		"",
		false, //  auto-ack disabled
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 Worker is waiting for messages...")

	for d := range msgs {
		log.Printf("📦 Processing: %s", d.Body)

		// simulate slow processing
		time.Sleep(2 * time.Second)

		// ⭐ success → acknowledge
		if string(d.Body) != "fail" {
			d.Ack(false)
			log.Println("✔ ACK sent")
		} else {
			// ⭐ failure → send back to queue
			d.Nack(false, true)
			log.Println("❌ NACK sent — message requeued")
		}
	}
}
