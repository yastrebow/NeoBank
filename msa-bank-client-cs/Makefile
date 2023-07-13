.PHONY: run migrateup migratedown
run:
	echo "Starting msa-bank-client-cs in port 8086" && \
	go run ./cmd/msa-bank-client-cs-server/ --port 8086
migrateup:
	migrate -path pkg/db/migrations/ -database "postgres://postgres:postgres@192.168.99.101:5438/url_db?sslmode=disable" -verbose up
migratedown:
	migrate -path pkg/db/migrations/ -database "postgres://postgres:postgres@192.168.99.101:5438/url_db?sslmode=disable" -verbose down
