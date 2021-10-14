test:
	go test -v -cover ./...
migrateup:
	go run ./db/migration/migrate.go up 1.0.0
migratedown:
	go run ./db/migration/migrate.go down 1.0.0
.PHONY: test