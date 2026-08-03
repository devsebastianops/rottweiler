ARG BUILD_TIMESTAMP
ARG BUILD_COMMIT
ARG BUILD_VERSION

FROM golang:1.26-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o bin/rottweiler -ldflags "-X github.com/devsebastianops/rottweiler/internal/cli.Version=${BUILD_VERSION} \
    -X github.com/devsebastianops/rottweiler/internal/cli.Commit=${BUILD_COMMIT} \
    -X github.com/devsebastianops/rottweiler/internal/cli.BuildTime=${BUILD_TIMESTAMP}" \
    ./cmd/rottweiler/main.go

FROM alpine:latest
COPY --from=build /app/bin/rottweiler /usr/bin/rottweiler
ENTRYPOINT ["/usr/bin/rottweiler"]