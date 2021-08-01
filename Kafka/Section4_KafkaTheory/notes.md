
### 8. Topics, paritions and offsets
- Topics: A particular stream of data. 
    - Similar to a table in a database. You can have as many topics as you want
- Topics are split in partitions (partitiion 0, 1, 3, n)
- Each partition is ordered
- Each message within a partition gets an incremental id called offset

#### Topic example: truck gps
- Comapany with many trucks and each truck has a gps
- Each truck reports it's gps to kafka
- Topic: trucks_gps
- Each truck will send a message to 20 seconds (imagine so)
- A topic is created with 10 partitions (arbitrary)
  
#### Topics, paritions and offsets part 2
- Offsets only have meaning for a specific partition
- Order is guaranteed only within a partition (not across partitions)
- Data is kept for a limited time (default is one week)
- Once a data is written to a partition, it can't be changed (immutability)
- Data is assigned to partitions randomly unless a key is provided 

### 9. Brokers and topics
- A Kafka cluster is composed of multiple brokers (servers)
- Each broker is identified with its ID (integer)
- Each broker contains certain topic partitions 
- After connecting to any broker (called a bootstrap broker) you will be connected to the entire cluster
- A good number to get started is 3. Large clusters can have over 100 brokers
- When you create a topic with multiple partitions, Kafka will distribute the partitions across brokers

### 10. Topic Replication Factor
- Topics should have a replication factor of > 1 (usually between 2 and 3)
- This way if a broker is down, another broker can serve the data
- At one time, only one broker can be a leder for a given partition, only that leader can receive and send data for a partition. The other brokers will synchronize the data
- One partition will have a leader and multiple ISR (in-sync replica)

### 11. Producers and message keys
- Producers write data to topics (which is made of partitions)
- Producers automatically know which broker and partition to write to
- In case of broken failures, producers recover
- Producers can choose to receive acknowledgement of data writes

#### Producers: Message keys
- Producers can choose to send a key with the message (string, number, etc)
- If there's no key, the message will be sent round robin
- If there is a key, all messages for that key will always go to the same partition
- A key is sent if you need message ordering for a specific field (ex: truck id). Key to partition is called hashing, it's a bit advanced

### 12. Consumers and consumer groups
- Consumers read data from a topic (identified by a name)
- They know hich broker to read from. In case of failures, they know how to recover
- Data is read in order within partitions

#### Consumer groups
- Consumer read data in consumer groups
- Each consumer from a consumer group reads from exclusive partitions
- If you have more consumers than partitions, some consumers will be inactive

### 13. Consumer offsets and Data Delivery semantics
- Kafka stores the offesets at which a consumer group has been reading
- Offsets are commited as the consumer group reads data and stored under __consumer_offsets
- When a consumer has processed data from kafka, it should be commiting the offsets
- If a consumer dies, it'll read back from where it left of

#### Delivery semantics
- You can choose when to commit your offsets
  - At most once: When the message is received
  - At least once (usually preferred): When the message is processed
  - Exactly once: Only available from kafka-to-kafka. For external systens, you need to use an idempotent consumer
  
### 14. Kafka Broker Discovery
- Every kafka broker is called a bootstrap server. You only need to connect to one broker and you'll be connected to the entire server.
- Each broker knows about all the brokers, topics and partitions
- Kafka client does a metadata reqeust, a broker replies with all the cluster info. Noice!

### 15. Zookeper
- Manages brokers (keeps a list of them)
- Zookeper helps in performing leader elections partitions. Sends notifications to kafka in case of changes
- Kafka can't work without zookeper (Dani says: This is no longer a requirement apparently)
- Zookeeper operates with an odd number of servers. Zookeper has a leader (handles writes) and the rest of the servers are followers (handle reads)
- Zookeper does not store consumer offsets 

### Quiz:
- Two consumers that have the same group.id (consumer group id) will read from mutually exclusive partitions (this is true)
