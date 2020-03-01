FROM debian:latest
MAINTAINER Dani
RUN apt-get update && apt-get install salt-master vim -y

EXPOSE 4505 4506 8000
CMD ["/usr/bin/salt-master", "-c", "/etc/salt"]

