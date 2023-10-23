run:
	go run main.go

docker-build:
	docker build -t effective_mobile_task:local .

docker-compose-up:
	docker compose -f docker-compose.yml up

