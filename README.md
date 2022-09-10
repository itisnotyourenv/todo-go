# Notes about this project

### Requirements

Работа с файлами конфигурации осуществляется с помощью библиотеки https://github.com/spf13/viper

Работа с миграциями https://github.com/golang-migrate/migrate

Работа с базой данных https://github.com/jmoiron/sqlx

Работа с Postgres https://github.com/lib/pq

Work with envrioment varialbes https://github.com/joho/godotenv 

#### Quick install golang-migrate

```bash
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
apt-get update
apt-get install -y migrate
```

### Work with golang-migrate CLI

#### Create migrate

```shell
migrate create -ext sql -dir ./schema -seq init
```



### Shortcuts

```shell
# running pg container in Docker
make pg
# start migrate
make migrate-up
# checking database tables
make pg-exec
# command for check tables with psql utilite
\d
```

***

### SetUp Porject steps.

#### Start Postgres container in Docker

```shell
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5431:5432 -d --rm postgres:13.0-alpine
```

#### Run migrates with golang-migrate CLI

```shell
migrate -path ./schema -database "postgresql://postgres:postgres@localhost:5431/postgres?sslmode=disable" up
```

