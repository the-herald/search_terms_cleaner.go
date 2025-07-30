# Use the official Go image
FROM golang:1.21 as builder

WORKDIR /app

# Copy go.mod and go.sum first (for layer caching)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Generate protobuf Go code
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN buf generate ./protos

# Build the Go app
WORKDIR /app/api
RUN go build -o /search-terms-cleaner

# ---

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /search-terms-cleaner /
ENTRYPOINT ["/search-terms-cleaner"]
