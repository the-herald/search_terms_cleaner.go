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

# Install Buf and required protoc plugins
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate code
RUN buf generate

# Build the Go app
WORKDIR /app/api
RUN go build -o /search-terms-cleaner

# Final minimal image
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /search-terms-cleaner /
ENTRYPOINT ["/search-terms-cleaner"]
