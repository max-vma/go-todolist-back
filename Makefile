include .env


MIGRATIONS_PATH = migrations
DB_URL = $(DATABASE_URL)

db_version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version
	
db_migrate:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

db_rollback:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

db_force:
	@echo "Forcing migration version to $${to:-0}..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $${to:-0}
	
db_migration:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make migration <name>"; \
	else \
		migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filter-out $@,$(MAKECMDGOALS)); \
	fi

%:
	@:
