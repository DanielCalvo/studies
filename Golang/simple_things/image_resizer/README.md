### Building on Linux

##### For windows:
```bash
GOOS=windows GOARCH=amd64 go build -o image_resizer.exe image_resizer.go
```

##### For Linux
```bash
go build image_resizer.go 
```

##### For Mac
```
GOOS=darwin GOARCH=amd64 go build image_resizer.go
```

Built with `go version go1.13.4 linux/amd64`

