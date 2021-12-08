build:
	docker-compose build avito_tech

run:
	docker-compose up avito_tech

app_down:
	docker-compose down avito_tech

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:49155/postgres?sslmode=disable' up

db_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:49155/postgres?sslmode=disable' down