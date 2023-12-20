MOCKERY_BIN := $(GOPATH)/bin/mockery
SWAGGER_BIN := $(GOPATH)/bin/swag

.PHONY: serve tidy test mock

air: 
	alias air='$(go env GOPATH)/bin/air'

docs:
	swag init -g /cmd/api/main.go -o docs

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