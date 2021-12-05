- To build and run:
```shell
docker build . -t docker_cmd_env_vars
docker run -e name=Bob -e surname=McBobson -e age=10 docker_cmd_env_vars
```