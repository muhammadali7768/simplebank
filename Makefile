postgres:
	docker run --name postgres17 -p 5432:5432 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine
createdb:
	docker exec -it postgres17 createdb --user=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank 
migrationup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrationup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrationdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrationdown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/muhammadali7768/simplebank/db/sqlc Store
test:
	go test -v -cover ./...
.PHONY: createdb dropdb postgres migrationup migrationdown migrationup1 migrationdown1 test server mock