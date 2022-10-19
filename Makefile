test:
	@go test -cover ./... -coverprofile cover.out
	@echo "-------------------------------------------------------------------------------------"
	@go tool cover -func cover.out
	@echo "-------------------------------------------------------------------------------------"

lint:
	@golangci-lint run

vendor:
	go mod vendor

dockerup:
	docker-compose up -d --build

dockerstop:
	docker-compose stop

migrate-create:
	migrate create -ext sql -dir ./migrate -seq init

migrate-up:
	migrate -path ./migrations -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' down

#db-force:
#	migrate -path ./migrations -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' force 1