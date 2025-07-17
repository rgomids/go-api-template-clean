# Production Dockerfile
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/api ./cmd/main.go

FROM alpine:latest
COPY --from=build /bin/api /bin/api
CMD ["/bin/api"]
