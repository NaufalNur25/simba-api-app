# Gunakan base image Go
FROM golang:3.21-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod download
RUN go mod tidy && go build -o ./api

EXPOSE 8080

CMD ["./api"]