# local
.PHONY: tidy
vendor:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

#dev
.PHONY: run
run:
	docker-compose run --rm prod

#dev
.PHONY: dev-fmt
dev-fmt:
	docker-compose run --rm dev go fmt ./...

.PHONY: dev-run
dev-run:
	docker-compose run --rm dev go run ./cmd/main.go

.PHONY: dev-infra
dev-infra:
	docker-compose run --rm database

.PHONY: dev-mi
dev-mi:
	docker-compose run --rm migrator goose up

.PHONY: dev-migrator
dev-migrator:
	docker-compose run --rm migrator

#test
.PHONY: test-infra
test-infra:
	docker-compose run --rm database-test

.PHONY: test-mi
test-mi:
	docker-compose run --rm migrator-test goose up

.PHONY: test-test
test-test:
	docker-compose run --rm test

# utilities
.PHONY: proto
proto:
	docker-compose run --rm dev sh scripts/protoc.sh

.PHONY: build
build:
	docker-compose build