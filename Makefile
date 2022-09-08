run:
	go run cmd/main.go

pg:
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5431:5432 -d --rm postgres:13.0-alpine
pgs:
	docker stop postgres
pgr:
	docker stop postgres
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5431:5432 -d --rm postgres:13.0-alpine
	docker logs -f postgres
pg-exec:
	docker exec -it postgres psql -U postgres

# example: make migrate-create NAME=third
migrate-create:
	migrate create -ext sql -dir schema -seq $(NAME)
migrate-up:
	migrate -path ./schema -database "postgresql://postgres:postgres@localhost:5431/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./schema -database "postgresql://postgres:postgres@localhost:5431/postgres?sslmode=disable" down
