MOCKERY_BIN := $(GOPATH)/bin/mockery

.PHONY: serve tidy test mock

alias:
	alias air='$(go env GOPATH)/bin/air' 
	alias swag='$(go env GOPATH)/bin/swag'

swag:
	swag init -g cmd/api/server/server.go -o docs

serve:
	go run cmd/api/main.go
tidy:
	go mod tidy && go mod vendor
test:
	go run cmd/test/main.go
mock:
	@echo "Generating mocks for interface $(interface) in directory $(dir)..."
	@$(MOCKERY_BIN) --name=$(interface) --dir=$(dir) --output=./internal/mocks
	cd ./internal/mocks && \
	mv $(interface).go $(filename).go
mig-up:
	go run cmd/migration/main.go -up
mig-down:
	go run cmd/migration/main.go -down
coverage:
	go test -v ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
seed:
	go run cmd/seed/main.go
mig-create:
	migrate create -ext sql -dir cmd/migration/migrations -seq $(name)