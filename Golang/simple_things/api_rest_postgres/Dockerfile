
FROM golang:1.15 AS builder
RUN mkdir /builddir
WORKDIR /builddir
#COPY go.mod .
#RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /builddir/app .

#Using ubuntu here just to debug, for a compiled app we could use a FROM scratch image
#FROM ubuntu
#COPY --from=builder /builddir/app .
#COPY --from=builder /builddir/entrypoint.sh .
#COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=chamber chamber .
#
#RUN chmod 755 entrypoint.sh chamber
#
#ENTRYPOINT ["./entrypoint.sh"]

FROM scratch
COPY --from=builder /builddir/app ./
EXPOSE "8080"
ENTRYPOINT ["./app"]