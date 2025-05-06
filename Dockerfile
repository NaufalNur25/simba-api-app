# Gunakan base image Go
FROM golang:latest AS builder

# Tentukan direktori kerja
WORKDIR /app

# Salin file Go yang diperlukan ke dalam image
COPY . .

# Bangun aplikasi Go
RUN go mod tidy && go build -o api .

# Tentukan port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./api"]