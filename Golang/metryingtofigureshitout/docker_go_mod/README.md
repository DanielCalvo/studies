- Using go modules: https://blog.golang.org/using-go-modules
- Very instructive: https://petomalina.medium.com/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888

- To build and run:
```shell
docker build . -t docker_go_mod && docker run docker_go_mod
```

- Module setup:
```shell
go mod init example.com/hello
```