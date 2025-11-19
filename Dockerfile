FROM golang:1.23-alpine AS development

WORKDIR /app

RUN go install github.com/air-verse/air@v1.52.3 && \
    go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
