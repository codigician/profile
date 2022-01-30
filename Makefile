up:
	docker build --progress=plain -t profile-api .
	docker run profile-api

run:
	go run .

unit-test:
	go test ./... -v -short

test:
	go test ./... -v

code-coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out | grep total | awk '{print $3}'

lint:
	golangci-lint run

mockgen:
	mockgen -destination=internal/mocks/about/mock_repository.go -package mocks -source=internal/about/service.go
	mockgen -destination=internal/mocks/mocks_profile_handler.go -package mocks -source=internal/handler.go