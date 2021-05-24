On .bashrc set this env var
- `export GO111MODULE=on`

Then:
- `go get github.com/spf13/cobra/cobra`
- `cobra init --pkg-name example` (make sure $PATH is setup to include $GOPATH/bin)
- `cd example`
- `go mod init example`
- `go get`

Enable _Go Modules Integration_ on Goland and set environment to `GO111MODULE=on`
