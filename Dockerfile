# Use the official Go image
FROM golang:1.24 as builder

WORKDIR /app

# Copy everything (early) so local replacements are available
COPY . .

# Download modules
RUN go mod tidy
RUN go mod download

# Build protobuf Go files
WORKDIR /app/protos
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN buf generate

# Build the Go app
WORKDIR /app/api
RUN go build -o /search-terms-cleaner

# Final minimal image
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /search-terms-cleaner /
ENTRYPOINT ["/search-terms-cleaner"]
