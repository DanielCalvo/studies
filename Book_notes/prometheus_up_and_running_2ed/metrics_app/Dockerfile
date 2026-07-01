FROM golang:alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /metrics-app .

FROM scratch

COPY --from=build /metrics-app /metrics-app
EXPOSE 8000 8001
ENTRYPOINT ["/metrics-app"]

