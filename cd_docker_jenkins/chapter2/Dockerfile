FROM ubuntu
MAINTAINER Dani
ENV NAME Banana
RUN apt-get update && \
apt-get install -y python
COPY hello.py .
ENTRYPOINT ["python", "hello.py"]

#docker build -t hello_world_python .
#docker run -e NAME=dani_pythontest hello_world_python