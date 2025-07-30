# Use the official Golang image
FROM golang:1.22

# Install Buf CLI
RUN curl -sSL \
  "https://github.com/bufbuild/buf/releases/download/v1.27.1/buf-Linux-x86_64.tar.gz" \
  | tar -xz -C /usr/local/bin --strip-components 1

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker caching
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Generate Go protobuf files using Buf
WORKDIR /app/Search_Terms_Cleaner/protos
RUN buf generate

# Return to main app root and build the binary
WORKDIR /app/Search_Terms_Cleaner
RUN go build -o server ./api/main.go

# Use a small final image for production
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=0 /app/Search_Terms_Cleaner/server .

# Required by Render
ENV PORT=10000

EXPOSE 10000
CMD ["/app/server"]
