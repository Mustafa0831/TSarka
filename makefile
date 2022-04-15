run:
	go run ./cmd/api/main.go -port 8080

test:
	go test ./... -v -cover -race

swag:
	swag init -d internal/api -g api.go --parseDependency --parseInternal

gen:
	mockgen -source=internal/counter/service.go -destination=./internal/counter/mocks/mock_counter_service.go
	mockgen -source=internal/counter/repository.go -destination=./internal/counter/mocks/mock_counter_repository.go
	mockgen -source=internal/hasher/service.go -destination=./internal/hasher/mocks/mock_hasher_service.go
	mockgen -source=internal/hasher/repository.go -destination=./internal/hasher/mocks/mock_hasher_repository.go
	mockgen -source=internal/email/service.go -destination=./internal/email/mocks/mock_email_service.go
	mockgen -source=internal/email/repository.go -destination=./internal/email/mocks/mock_email_repository.go
	mockgen -source=internal/firstname/service.go -destination=./internal/firstname/mocks/mock_firstname_service.go
	mockgen -source=internal/firstname/repository.go -destination=./internal/firstname/mocks/mock_firstname_repository.go


.PHONY: run test swag gen