FROM golang:1.10 AS builder

ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/bin/dep

RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/DanielCalvo/unnamed-exporter
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE "2113"
ENTRYPOINT ["./app"]