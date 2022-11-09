postgres:
	docker run --name postgres-docker -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it postgres-docker createdb --username=root --owner=root sample_bank
dropdb:
	docker exec -it postgres-docker dropdb sample_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/sample_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY:
	postgres createdb dropdb, migrateup, migratedown, sqlc


