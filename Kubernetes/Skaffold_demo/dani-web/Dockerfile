FROM golang:1.13.4-alpine3.10 as builder
COPY web.go .
RUN go build -o /web .

FROM alpine:3.10
CMD ["./web"]
COPY --from=builder /web .