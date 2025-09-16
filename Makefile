server-up:
	cd cmd/server && go run main.go

docker-up:
	docker compose -f docker.compose.yml up -d