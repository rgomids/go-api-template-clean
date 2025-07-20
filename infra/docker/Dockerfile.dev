# Development Dockerfile
FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go install github.com/cespare/reflex@latest
CMD ["go", "run", "./cmd/main.go"]
