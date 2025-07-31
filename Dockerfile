# === Stage 1: Builder ===
FROM golang:1.24 as builder

WORKDIR /app

# Copy mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Generate protobuf Go code
RUN go install github.com/bufbuild/buf/cmd/buf@latest
WORKDIR /app/protos
RUN buf generate .

# Build the Go app
WORKDIR /app/api
RUN go build -o /search-terms-cleaner

# === Stage 2: Runtime ===
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /search-terms-cleaner /
ENTRYPOINT ["/search-terms-cleaner"]
