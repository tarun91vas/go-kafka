## Kafka setup and scritps

### Kafka Installtion

```
./install.sh /path/to/install
```


### Quick start using console

- Edit config/server.properties: `advertised.listeners=PLAINTEXT://localhost:9092` 
- Start zookeeper: `bin/zookeeper-server-start.sh config/zookeeper.properties`
- Start kafka server: `bin/kafka-server-start.sh config/server.properties`
- Create topic: `bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test`
- List topics on zookeeper: `bin/kafka-topics.sh --list --zookeeper localhost:2181`
- Start producer: `bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test`
- Start consumer: `bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning`
- Send messages in producer console to receive at consumer console
- For a multi broker cluster, start more than one broker by creating more server.properties file and increasing the replication factor as described [here](https://kafka.apache.org/quickstart#quickstart_multibroker)

## Kafka producer and consumers in go

### Prerequisites
- [gb tool](https://getgb.io/) 
- [gb-vendor](https://godoc.org/github.com/constabulary/gb/cmd/gb-vendor) plugin

### Dependencies

To install vendor dependencies
```
./dependencies.sh
```

### Run
- Producer

```
$ gb build kafka-producer
$ ./bin/kafka-producer
```

- Consumer

```
$ gb build kafka-consumer
$ ./bin/kafka-consumer
```




