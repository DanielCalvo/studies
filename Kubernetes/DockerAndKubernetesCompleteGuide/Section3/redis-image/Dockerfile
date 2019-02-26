#Use an existing Docker image as a base
FROM alpine

#Download and install a dependency
RUN apk add --update redis

#What to do when it starts as a container
CMD ["redis-server"]