FROM golang:1.23-alpine AS development

WORKDIR /app

# Install air and swag
RUN go install github.com/air-verse/air@v1.52.3 && \
    go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN swag init

EXPOSE 8080

# Use air to watch and auto-reload
CMD ["air", "-c", ".air.toml"]
