# запуск тестов
test:
	go test -cover ./...

lint:
	@golangci-lint run

vendor:
	go mod vendor

bench:
	go test -bench=BenchmarkExecute -benchmem -benchtime 5s -count=5

migrate-create:
	migrate create -ext sql -dir ./migrate -seq init

migrate-up:
	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' up

migrate-down:
	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' down

#db-force:
#	migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' force 1