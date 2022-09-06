# Notes about this project

Работа с файлами конфигурации осуществляется с помощью библиотеки https://github.com/spf13/viper

Работа с миграциями https://github.com/golang-migrate/migrate

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



### Startup steps

```shell
# running pg container in Docker
make pg
# start migrate
make migrate
# checking database tables
make pg-exec
# command for check tables with psql utilite
\d
```

