.PHONY: generate build run doc validate spec clean help create-file-migration

all: clean spec generate build

validate:
	swagger validate ./api/biofarma/index.yml

spec:
	swagger flatten ./api/biofarma/index.yml -o=./api/biofarma/result.yml --format=yaml --with-flatten=full
	swagger flatten ./api/bing_maps/index.yml -o=./api/bing_maps/result.yml --format=yaml --with-flatten=full

build: 
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/cli
	
run-binary:
	./cli api --port=8080 --host=0.0.0.0

run-local:
	go run cmd/cli/main.go api --port=8080 --host=0.0.0.0

doc: validate
	swagger serve api/biofarma/index.yml --no-open --host=0.0.0.0 --port=8080 --base-path=/

clean:
	rm -rf cli
	rm -rf ./gen/client
	rm -rf ./gen/models
	rm -rf ./gen/restapi
	go clean -i .

generate: validate
	swagger generate server --exclude-main -A server -t gen -f ./api/biofarma/result.yml --principal models.Principal
	swagger generate client -A bing-maps -c gen/client/bing_maps_client -m gen/models -f ./api/bing_maps/result.yml

migrate-create-file:
	go run cmd/cli/main.go migration create $(Arguments)

migrate-up:
	go run cmd/cli/main.go migration up

migrate-down:
	go run cmd/cli/main.go migration down

migrate-force:
	go run cmd/cli/main.go migration force

help:
	@echo "make: compile packages and dependencies"
	@echo "make validate: OpenAPI validation"
	@echo "make spec: OpenAPI Spec"
	@echo "make clean: remove object files and cached files"
	@echo "make build: Generate Server and Client API"
	@echo "make doc: Serve the Doc UI"
	@echo "make run: Serve binary file"
	@echo "make run-local: Serve main.go"
	@echo "make migrate-create-file: Create new file migration"
	@echo "make migrate-up: Migrate up"
	@echo "make migrate-down: Migrate down"
	@echo "make migrate-force: Fix dirty migration"

Arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
