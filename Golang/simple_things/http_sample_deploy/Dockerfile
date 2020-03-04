FROM golang:1.13 AS builder

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE "2113"
ENTRYPOINT ["./app"]