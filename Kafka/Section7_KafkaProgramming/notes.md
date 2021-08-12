```shell
#Start kafka
docker-compose up
kafka-topics.sh --bootstrap-server localhost:9092 --topic myTopic --create

#(...)

docker-compose down
```