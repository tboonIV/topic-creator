build docker image

```sh
docker build -t topic-creator:1.0.0 .
```

Required env configs:
- BOOTSRAP_SERVER: bootstrap server host and port. eg localhost:9092
- TOPIC_NAMES: topic names in string with coma separator. eg my-topic,his-topic