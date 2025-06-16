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

# Собираем приложение и утилиты
RUN make build && \
    go install github.com/pressly/goose/v3/cmd/goose@v3.20.0 && \
    cp /go/bin/goose /app/bin/goose

# Финальная стадия
FROM alpine:3.21

WORKDIR /app

# Устанавливаем bash для выполнения скриптов
RUN apk add --no-cache bash

# Копируем бинарники и миграции
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/local.env /app/local.env


EXPOSE 8080

CMD ["make all"]
