package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var brokers = []string{"jorgekafkaserver.westeurope.cloudapp.azure.com:9092"}

func main() {

	//setup relevant config info
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	config := sarama.NewConfig()
	config.ClientID = "jorgeapp"
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, nil)

	if err != nil {
		log.Fatal(err)
	}

	topic := "test"                              //e.g create-user-topic
	var partition int32 = 1                      //Partition to produce to
	msg := "actual information to save on kafka" //e.g {"name":"John Doe", "email":"john.doe@email.com"}
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Value:     sarama.StringEncoder(msg),
	}
	partition, offset, err := producer.SendMessage(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("offset: %s, partition: %s", fmt.Sprint(offset), fmt.Sprint(partition))
}
