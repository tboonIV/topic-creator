package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func validate(vars []string) {
	for _, varName := range vars {
		if os.Getenv(varName) == "" {
			log.Fatal(varName + " environment variable is not set")
		}
	}
}

func main() {
	// Validate required environment variables
	validate([]string{"BOOTSTRAP_SERVERS", "TOPIC_NAMES"})

	// Read an environment variable
	BOOTSTRAP_SERVERS := os.Getenv("BOOTSTRAP_SERVERS")
	fmt.Printf("BOOTSTRAP_SERVERS value: %s\n", BOOTSTRAP_SERVERS)
	TOPIC_NAMES := os.Getenv("TOPIC_NAMES")
	fmt.Printf("TOPIC_NAMES value: %s\n", TOPIC_NAMES)

	// Create a new Kafka administration client
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": BOOTSTRAP_SERVERS,
	})
	if err != nil {
		log.Fatalf("Failed to create admin client: %v", err)
	}
	defer adminClient.Close()

	// Create the topic
	fmt.Println("Creating Kafka topics " + TOPIC_NAMES)
	topicNameList := strings.Split(TOPIC_NAMES, ",")
	for _, topicName := range topicNameList {
		// Create a new topic specification
		topicSpec := kafka.TopicSpecification{
			Topic:             topicName,
			NumPartitions:     1,
			ReplicationFactor: 1,
		}

		result, err := adminClient.CreateTopics(context.Background(), []kafka.TopicSpecification{topicSpec})
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		// Wait for the result to complete
		for _, topicResult := range result {
			if topicResult.Error.Code() != kafka.ErrNoError {
				log.Fatalf("Failed to create topic %s: %v", topicResult.Topic, topicResult.Error)
			}
			fmt.Printf("Topic %s created successfully\n", topicResult.Topic)
		}
	}
}
