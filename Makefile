DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres17 -p 5432:5432 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine
createdb:
	docker exec -it postgres17 createdb --user=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank 
migrationup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrationup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migrationdown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrationdown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1


sqlc:
	sqlc generate

server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/muhammadali7768/simplebank/db/sqlc Store
test:
	go test -v -cover ./...
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb \
    --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl
.PHONY: createdb dropdb postgres migrationup migrationdown migrationup1 migrationdown1 test server mock proto evans