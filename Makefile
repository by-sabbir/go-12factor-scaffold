.PHONY:

build:
	docker compose build

run:
	docker compose up -d

migrate:
	docker compose exec app sh -c "./srv migrate up"

test:
	go test -cover ./...