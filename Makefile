test:
	go test -cover ./...

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
	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' up

migrate-down:
	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' down

#db-force:
#	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' force 1