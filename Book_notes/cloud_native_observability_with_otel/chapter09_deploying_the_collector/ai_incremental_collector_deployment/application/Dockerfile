# Build a static Go binary so the final image needs no operating-system files.
FROM golang:1.26.2-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -o /chapter9-greeting .

FROM scratch
COPY --from=build /chapter9-greeting /chapter9-greeting
EXPOSE 8080
ENTRYPOINT ["/chapter9-greeting"]
