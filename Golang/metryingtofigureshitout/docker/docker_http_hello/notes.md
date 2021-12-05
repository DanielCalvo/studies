- To build and run:
```shell
docker build . -t docker_http_hello
docker run -p 8080:8080 -e name=Bob -e surname=McBobson -e age=10  docker_http_hello
```