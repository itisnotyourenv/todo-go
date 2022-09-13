run:
	go run cmd/main.go

# need set DB_PASS to os env
# example: export DB_PASS=123456
pg:
	docker run --name postgres -e POSTGRES_PASSWORD=${DB_PASS} -p 5431:5432 -d --rm postgres:13.0-alpine
pgs:
	docker stop postgres
pgr:
	docker stop postgres
	docker run --name postgres -e POSTGRES_PASSWORD=${DB_PASS} -p 5431:5432 -d --rm postgres:13.0-alpine
	docker logs -f postgres
pg-exec:
	docker exec -it postgres psql -U postgres

# example: make migrate-create NAME=third
migrate-create:
	migrate create -ext sql -dir schema -seq $(NAME)
# example: make migrate-up DB_PASS=123456
migrate-up:
	migrate -path ./schema -database "postgresql://postgres:${DB_PASS}@localhost:5431/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./schema -database "postgresql://postgres:${DB_PASS}@localhost:5431/postgres?sslmode=disable" down
