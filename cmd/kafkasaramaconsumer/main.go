package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var brokers = []string{"jorgekafkaserver.westeurope.cloudapp.azure.com:9092"}

func main() {

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	config2 := sarama.NewConfig()
	//config2.Group.Mode = cluster.ConsumerModePartitions
	config2.Consumer.Offsets.Initial = sarama.OffsetOldest
	config2.Consumer.Return.Errors = true
	//config2.Group.Return.Notifications = true

	consumer, err := sarama.NewConsumer(brokers, config2)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, 1)

	if err != nil {
		log.Fatal(err)
	}

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d  %s\n", msg.Offset, msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)

}
