server-up:
	go run ./cmd/server

docker-up:
	docker compose -f docker.compose.yml up -d