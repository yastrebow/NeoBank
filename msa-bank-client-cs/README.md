Для создания схем, таблиц и индексов в БД необходимо выполнить "migrate -database postgres://{{DB_URL}}?sslmode=disable"&"x-migrations-table=msa_bank_client_cs -path migrations -verbose up"
Для генерации кода из Swagger файла "swagger generate server -f ./api/msa-bank-client-cs.yml -A msa-bank-client-cs"
Для запуска приложения  "go run .\cmd\msa-bank-client-cs-server\main.go"