# Билд стадия
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git make bash

# Копируем файлы модулей сначала для кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

RUN make build install-deps

# Финальная стадия
FROM alpine:3.21

WORKDIR /app

# Устанавливаем bash для выполнения скриптов
RUN apk add --no-cache bash make

# Копируем бинарники и миграции
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/local.env /app/local.env
COPY --from=builder /app/Makefile /app/Makefile



EXPOSE 8080

CMD sh -c "make migration-up && /app/bin/main"