postgres:
	docker run --name postgres-docker-maker -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createDb:
	docker exec -it postgres-docker-maker createdb --username=root --owner=root sample_bank
dropDb:
	docker exec -it postgres-docker-maker dropdb sample_bank

migrateUp:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable" -verbose up
migrateDown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

# command for running go test cases
test: 
	go test -v -cover ./...

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc test


