include .env

MIGRATIONS_PATH = migrations
DB_URL = $(DATABASE_URL)

migrate:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-rollback:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
