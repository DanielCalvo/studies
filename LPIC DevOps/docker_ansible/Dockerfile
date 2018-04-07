FROM debian:latest
LABEL description="This is a simple debian dockerfile"

RUN apt-get update
RUN apt-get install gnupg -y
RUN echo "deb http://ppa.launchpad.net/ansible/ansible/ubuntu trusty main" >> /etc/apt/sources.list
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 93C4A3FD7BB9C367 #Requires gnupg!
RUN apt-get update
RUN apt-get install ansible -y


#Bringing up the instance
#cd docker_ansible
#docker build -t docker_ansible .
#docker run docker_ansible