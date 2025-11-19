FROM golang:1.21-alpine AS development

WORKDIR /app

# Install air for hot-reloading
RUN go install github.com/cosmtrek/air@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Run with air
CMD ["air", "-c", ".air.toml"]
