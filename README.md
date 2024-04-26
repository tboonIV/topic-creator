# Intro

Docker image that create Kafka Topics upon start up and exit when completed.

# Example usage 

docker-compose.yml

```yml
topic-creator:
  container_name: topic-creator
  image: tboon/topic-creator:1.0.0
  environment:
    BOOTSTRAP_SERVERS: kafka-0:9092
    TOPIC_NAMES: my-topic,his-topic
  depends_on:
    - kafka-0
```

# Required Configs

Following environment variables required for `topic-creator` docker start up:
- BOOTSRAP_SERVER: bootstrap server host and port. eg localhost:9092
- TOPIC_NAMES: topic names in string with coma separator. eg my-topic,his-topic
