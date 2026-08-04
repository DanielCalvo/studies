# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/image-resizer-api .

FROM scratch

USER 65532:65532

COPY --from=build --chown=65532:65532 /out/image-resizer-api /image-resizer-api

EXPOSE 8080

ENTRYPOINT ["/image-resizer-api"]
