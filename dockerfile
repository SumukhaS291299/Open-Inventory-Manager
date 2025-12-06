# ============================
# STAGE 1 — Build the Go binary
# ============================
FROM golang:1.22-alpine AS builder

# Install build tools (optional but recommended)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o inventory-server ./main.go

# --------------------------------------------------------------------------------------------------------------

# ============================
# STAGE 2 — Run the application
# ============================
FROM alpine:3.20

# Add certificates (so HTTPS and outside APIs work)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the built binary from builder stage
COPY --from=builder /app/inventory-server .

# Expose the port your Gin server uses (change if needed)
EXPOSE 8080

# Run the server
CMD ["./inventory-server"]
