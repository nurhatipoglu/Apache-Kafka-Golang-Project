package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {

		payload := struct {
			Message string `json:"message"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(payload.Message),
		}, nil)

		return c.SendString("Nur Selam 👋!")
	})

	app.Listen(":3000")

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
