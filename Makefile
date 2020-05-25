build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

run: build
	docker-compose up --build taxi_service
