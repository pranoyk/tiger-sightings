generate:
	@echo "Generating..."
	@go generate ./...
	@echo "Done"

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ./db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path ./db/migrations down