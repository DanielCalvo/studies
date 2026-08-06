# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine AS build
WORKDIR /src
COPY main.go ./
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GO111MODULE=off GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/traffic-gen .

FROM scratch
USER 65532:65532
COPY --from=build --chown=65532:65532 /out/traffic-gen /traffic-gen
COPY --chown=65532:65532 test-data /test-data
ENTRYPOINT ["/traffic-gen"]
