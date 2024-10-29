FROM cgr.dev/chainguard/go:latest AS builder
ARG TARGETOS TARGETARCH

WORKDIR /app
# dependencies, add local, dependant package here
COPY protocol/ protocol/
COPY sdk/ sdk/
COPY lib/ocrypto lib/ocrypto
COPY lib/flattening lib/flattening
COPY lib/fixtures lib/fixtures
COPY service/ service/
COPY examples/ examples/
COPY go.work go.work.sum ./
RUN cd service \
    && go mod download \
    && go mod verify
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o opentdf ./service

FROM cgr.dev/chainguard/glibc-dynamic

# Copy the opentdf binary
COPY --from=builder /app/opentdf /usr/bin/

# Set the working directory to /app, so the paths match the expected structure
WORKDIR /app

ENTRYPOINT ["/usr/bin/opentdf"]
