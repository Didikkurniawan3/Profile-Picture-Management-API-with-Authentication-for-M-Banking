FROM golang:alpine

ENV DB_NAME     postgres
ENV DB_USER     postgres
ENV DB_HOST     localhost
ENV DB_PASSWORD postgres
ENV DB_PORT     5432
ENV STAGE       development

# Buat direktori kerja
WORKDIR /app

# Salin file go.sum dan go.mod, lakukan pengunduhan dependensi
COPY go.* ./
RUN go mod download

# Salin seluruh kode aplikasi
COPY . .

# Kompilasi dan jalankan aplikasi
CMD go run main.go
