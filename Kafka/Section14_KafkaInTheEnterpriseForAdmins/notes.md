### 99. Kafka Cluster Setup High Level Architecture Overview
- You want at least 3 servers (old zookeeper) in 3 different AZs (for AWS)

#### A few gotchas
- It's not easy to set up a cluster!
- Isolate zookeper and broker on separate servers (no longer applicable?)
- Monitoring needs to be implemented
- Operations have to be mastered
- You need a good kafka admin! 
- Use a managed kafka as a service if you can!

### 100. Kafka Monitoring & Operations
- Kafka exposes metrics through JMX
- These metrics are important monitoring kafka and making sure it runs correctly
- Common places to host kafka metrics:
    - Datadog, ELK, Prometheus, etc

- Relevant metrics:
    - Under replicated partitions: Partitions having trouble with the in-sync replicas (could be high load)
    - Request handlers: Overall usage of the kafka cluster
    - Request timing: How long it takes to reply requests
- Kafka, confluent and datadog documentations are recommended for monitoring info

#### Operations
- Kafka operations teams must be able to perform the following tasks
- Rolling restart of brokers
- Updating configurations
- Rebalancing partitions
- Increasing replication factor
- Adding a broker
- Replacing a broker
- Removing a broker
- Upgrading a kafka cluter with 0 downtime

Sounds like a fun challenge to learn!

### 101. Kafka Security
- When you first launch a cluster from local, any client can access your kafka cluster
- All data sent is visible on the network
- Someone could intercept data, publish bada data or delete topics

#### Encryption in Kafka!
- Kafka can use encryption between broker and clients, similar to HTTPs
- Encrypted data goes on port 9093
- Clients can authenticate to Kafka!
    - SSL auth over certs
- Authorization in Kafka
    - You can grant topic rights per user
    - There are ACLs for access
- Currently it is hard to set up Kafka security!

### 102. Kafka Multi Cluster & MirrorMaker
- Kafka can only operate well within a single reagion
- Global enterprise: Kafka clusters all around the world, with some level of replication!
- Mirror maker: Open source tool that ships with apache kafka for replication
    - Netflix uses Flink
    - Uber uses ureplicator, it's open sauce
    - Comcast has their own open source kafka connect store
    - Confluent has a paid thing
- There are many options you can try

- There are two designs for cluster
    - Active passive: One cluster receives from producers, but on
    - Active active: producers and consumers to both clusters, and mirror maker sorts it out
- The whole replication thing is an advanced topic you might not need and it's kinda hard
- Replicating does not preserve offsets, just data!