FROM golang:1.15

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /myapp

#COPY . $GOPATH/src/mypkg/myapp/
#WORKDIR $GOPATH/src/mypkg/myapp/
#
#RUN go get -d -v
#RUN go build -o /go/bin/myapp

ENTRYPOINT ["/myapp"]