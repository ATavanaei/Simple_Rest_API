package KafkaSet

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"fmt"
)

func KafkaConsume(server,port,topicSet string, return_Func func([]string)) {
	data_get := []string{}
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server + ":" + port,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topicSet, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			data_get = append(data_get,string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
	return_Func(data_get)
}

func kafkaProduce(server,port,topicSet string, data_set []string) (messageDeliver bool){

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server + ":" + port})
	if err != nil {
		panic(err)
	}
	defer p.Close()
	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					messageDeliver = false
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					messageDeliver = true
				}
			}
		}
	}()
	// Produce messages to topic (asynchronously)
	topic := topicSet
	for _, Message_set := range data_set {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(Message_set),
		}, nil)
	}
	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return
}