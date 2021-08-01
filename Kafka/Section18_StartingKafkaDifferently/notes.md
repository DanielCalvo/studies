```shell
export CONFLUENT_HOME="/home/daniel/kafka/confluent_home"
confluent local services kafka start
```

### 117. Start Kafka Development environment using Docker
```shell
cd /tmp
git clone git@github.com:conduktor/kafka-stack-docker-compose.git
cd kafka-stack-docker-compose
docker-compose -f zk-single-kafka-single.yml up

kafka-topics.sh --bootstrap-server localhost:9092 --topic my_topic --create
kafka-console-producer.sh --broker-list localhost:9092 --topic my_topic
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic my_topic --from-beginning
```

### 118. Starting a multi broker Kafka Cluster using Docker
```shell
#cd tmp, clone and cd into dir same as above
docker-compose -f zk-multiple-kafka-multiple.yml up

kafka-topics.sh --bootstrap-server localhost:9092 --topic my_topic_2 --partitions 12 --replication-factor 3 --create
kafka-topics.sh --bootstrap-server localhost:9092 --describe
kafka-console-producer.sh --broker-list localhost:9092,localhost:9093,localhost:9094 --topic my_topic_2
kafka-console-consumer.sh --bootstrap-server localhost:9093 --topic my_topic_2 --from-beginning #You can choose any of the brokers!
docker-compose -f zk-multiple-kafka-multiple.yml down
```

### 119. Kafka Advertised Host Setting
- Advertised listeners is the most important setting of Kafka!
- It's the IP/DNS name that kafka will listen on
- If you have kafka listening on localhost, it's fine if all is on the same machine
- If it's listening on a NAT address (192.168.x.x) you won't be able to connect to it through the internet
- If it's listening on a public address, be careful, the public address can change

#### In a Nutshell
- If your clients are on your private network
    - Set either the internal IP or the internal DNS
- If your clients are on a public network
    - Set either the external IP or the public DNS

### 120. Starting Kafka on a Remote Machine
- Stephane install kafka on a EC2 machine with a SG rule so that kafka can be accessed through the internet