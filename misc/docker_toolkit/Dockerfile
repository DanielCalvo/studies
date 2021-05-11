FROM ubuntu
RUN apt-get update && apt-get install vim git -y
RUN useradd -ms /bin/bash daniel
USER daniel
WORKDIR /home/daniel