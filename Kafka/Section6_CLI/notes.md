### 32. Kafka Topics CLI

```shell
kafka-topics.sh --bootstrap-server localhost:9092 --topic first_topic --create #This worked!

#In Kafka, you cannot create a topic with a number of replication-factors larger than the number of brokers you have
#What Stephane ran:
kafka-topics.sh --bootstrap-server localhost:9092 --topic first_topic --create --partitions 3 --replication-factor 1
kafka-topics.sh --bootstrap-server localhost:9092 --topic second_topic --create --partitions 6 --replication-factor 1

kafka-topics.sh --bootstrap-server localhost:9092 --list
kafka-topics.sh --bootstrap-server localhost:9092 --topic first_topic --describe #Noice!
```
 
### 33. Kafka Console Producer CLI
```shell
kafka-console-producer.sh --broker-list localhost:9092 --topic first_topic
#Enter messages!
kafka-console-producer.sh --broker-list localhost:9092 --topic first_topic --producer-property acks=all

kafka-console-producer.sh --broker-list localhost:9092 --topic new_topic
#Gives a warning at first, but then stops
#This new topic was created, producing to a topic that does not exsist gets it created
kafka-topics.sh --bootstrap-server localhost:9092 --list

#Stephane says: Always create a topic beforehand
kafka-topics.sh --bootstrap-server localhost:9092 --topic new_topic --describe #Noice!
```

### 34. Kafka Console Consumer CLI
```shell
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first_topic #Only newly produced messages will appear, not old ones
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first_topic --from-beginning #Cool!
```

### 35. Kafka Consumers in Group
```shell
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first_topic --group my-first-group
#If you start multiple consumers with the same application group, the consuming will be split between them
```

### 36. Kafka Consumer Groups
```shell
kafka-consumer-groups.sh #Lists, describes, deletes, or resets consumer groups offset
kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
kafka-console-producer.sh --broker-list localhost:9092 --topic first_topic --group my-first-group
```

### 37. Resetting offsets
- If you reset the offset for a given group, the data will be consumed again from the offset that you set up

### 38. CLI Options that are good to know
```shell
#Producer with keys
kafka-console-producer --broker-list 127.0.0.1:9092 --topic first_topic --property parse.key=true --property key.separator=,
> key,value

#Consumer with keys
kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic first_topic --from-beginning --property print.key=true --property key.separator=,
```

### 39. UI: Conduktor
- Check out Conduktor if you're interested
