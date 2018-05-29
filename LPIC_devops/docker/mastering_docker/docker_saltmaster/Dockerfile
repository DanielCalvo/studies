FROM debian:latest
MAINTAINER Dani
RUN apt-get update && apt-get install salt-master -y

RUN mkdir /myvol
RUN echo "hey guys!" > /myvol/hello
VOLUME /myvol

EXPOSE 4505 4506
CMD ["/usr/bin/salt-master", "-c", "/etc/salt"]

