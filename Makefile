football_database_service_api:
	go build -o ./bin/football-database-service-api cmd/footballdb-api/main.go
football_database_service:
	go build -o ./bin/football-database-service cmd/footballdb/main.go
