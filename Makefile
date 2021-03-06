test:
	go test -v -cover ./...
migrateup:
	go run ./db/migration/migrate.go up 1.0.0
migratedown:
	go run ./db/migration/migrate.go down 1.0.0
resetdb:
	go run ./db/migration/migrate.go down 1.0.0
	go run ./db/migration/migrate.go up 1.0.0
dev:
	cd cmd/http-server/ && go run *.go
.PHONY: test migrateup migratedown resetdb dev