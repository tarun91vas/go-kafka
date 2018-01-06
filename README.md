## Kafka setup and scritps

### Installtion

```
./install.sh /path/to/install
```


### Quick  start

- Start zookeeper: `bin/zookeeper-server-start.sh config/zookeeper.properties`
- Edit config/server.properties: `advertised.listeners=PLAINTEXT://localhost:9092`
- Start kafka server: `bin/kafka-server-start.sh config/server.properties`
- Create topic: `bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test`
- List topics on zookeeper: `bin/kafka-topics.sh --list --zookeeper localhost:2181`
- Start producer: `bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test`
- Start consumer: `bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning`
- Send messages in producer console to receive at consumer console
