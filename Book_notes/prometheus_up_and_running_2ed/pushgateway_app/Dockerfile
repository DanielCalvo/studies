FROM golang:alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /pushgateway-app .

FROM scratch

COPY --from=build /pushgateway-app /pushgateway-app
ENTRYPOINT ["/pushgateway-app"]

