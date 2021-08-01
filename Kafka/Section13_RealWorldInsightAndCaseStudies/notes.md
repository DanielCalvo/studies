### 91. Choosing Partition Count & Replication Factor
- The two most important parameters when creating a topic
- They impact performance and durability of the system overall
- It's best to get the parameters right the first time!
    - If the partition count increases during a topic lifecycle, you will break your keys ordering guarantees
    - If the replication factor increases during a topic lifecycle, you put more pressure on your cluster, which can lead to a decrease in performance

#### Partitions count
- Each partition can handle a throughput of a few MB/s (measure it for your set up)
- More partitions imply
    - Better parrllelism, better throughput
    - Ability to run more consumers in a group to scale
    - Ability to leverage more brokers if you have a large cluster
    - But: More elections for zoopeker (deprecated?) more open files in Kafka
    
#### Guidelines:
- Partitions per topic: Million dollar question!
    - Small cluster (< 6 brokers): Have the number of partitions be 2x the number of brokers (rough estimate)
    - Big cluster (> 12 brokers): One partition per broker
    - Adjust for consumers: If your consumer number is high (ex: 20 at the same time) then you probably want 20 partitions per topic, regardless of how big or small your cluster is
    - Adjust for producers: IF you know you have high throughput, do 3x partition per number of brokers so you can adjust in the future
- Test! Every kafka cluster will have different performance. 
- Don't go for extremes. Don't create a topic with 1000 partitions, that's useless! 

#### Replication Factor
- At least 2, usually 3, maximum 4
- The higher the replication factor
    - More resilient (N-1 brokers can fail)
    - But more replication (higher latency if acks=all)
    - More disk space on your system (50% more if RF is 3 instead of 2)

- Guideline: 3 to get started
    - If replication performance is an issue, get a better broker (better machine)
    - Never set it to 1 in production!

#### Cluster guidelines
- It is advised that a broker should not gold more than 2000 to 4000 partitions across all topics of that broker
- Additionally, an entire kafka cluster should have a maximum of 20.000 partitions across all brokers 
- If you need more partitions in your cluster, add brokers instead (?)
- If you need more than 20.000 partitions in your cluster, follow the netflix model and create more kafka clusters

### Case studies
- These were cool, but I'm still lacking some practice and solid KAfka knowledge to suggest architectures for those
