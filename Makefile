generate:
	@echo "Generating..."
	go generate ./...
	@echo "Done"

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ./db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path ./db/migrations down

migrate-force:
	migrate -path ./db/migrations -database ${POSTGRESQL_URL} force ${VERSION}

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down