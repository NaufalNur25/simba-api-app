# Gunakan base image Go
FROM FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy go mod dan go sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy semua file project ke dalam container
COPY . .

# Build binary-nya
RUN go build -o main .

# Expose port (disesuaikan dengan yang kamu pakai di Gin)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./api"]