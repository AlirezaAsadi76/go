createdb:
	sudo docker exec -it bank createdb --username=root --owner=root mybank
dropdb:
	sudo docker exec -it bank dropdb mybank
postgres:
	sudo docker run --name bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:14.4-alpine
migrateUp:
	migrate -path db/migration -database "postgres://root:123456@localhost:5432/mybank?sslmode=disable" -verbose up
migrateDown:
	migrate -path db/migration -database "postgres://root:123456@localhost:5432/mybank?sslmode=disable" -verbose down
.PHONY: createdb dropdb postgres migrateUp migrateDown
