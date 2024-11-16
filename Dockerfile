# Вказуємо базовий образ
FROM golang:1.23.3-alpine

# Створюємо робочу директорію
WORKDIR /app

# Копіюємо go.mod і go.sum
COPY go.mod go.sum ./

# Завантажуємо залежності
RUN go mod download

# Копіюємо решту файлів
COPY . .

# Будуємо додаток
RUN go build -o convert

# Вказуємо команду запуску
CMD ["./convert"]

# Вказуємо порт
EXPOSE 8080
