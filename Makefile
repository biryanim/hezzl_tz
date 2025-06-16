include local.env

LOCAL_BIN = $(CURDIR)/bin

all: install-deps migration-up build

build:
	go build -o bin/main ./cmd/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.20.0

migration-up:
	@echo "=== PostgreSQL ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/postgres postgres ${PG_MIGRATION_DSN} up -v
	@echo "\n=== ClickHouse ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/clickhouse clickhouse ${CH_MIGRATION_DSN} up -v



migration-down:
	@echo "=== PostgreSQL ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/postgres postgres ${PG_MIGRATION_DSN} down -v
	@echo "\n=== ClickHouse ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/clickhouse clickhouse ${CH_MIGRATION_DSN} down -v

migration-status:
	@echo "=== PostgreSQL ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/postgres postgres ${PG_MIGRATION_DSN} status -v
	@echo "\n=== ClickHouse ==="
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR}/clickhouse clickhouse ${CH_MIGRATION_DSN} status -v