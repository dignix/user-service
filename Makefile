.PHONY: build
build:
	go build -o ./bin/app main.go

.PHONY: run
run:
	./bin/app

.PHONY: build-and-run
build-and-run: build run

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: test-bench
test-bench:
	go test -bench=. ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.PHONY: format
format:
	gofmt -w ./..

.PHONY: gen
gen:
	protoc -I api/v1/pb --go_out=plugins=grpc:. --grpc-gateway_out=:. --swagger_out=:api/v1 user.proto

.PHONY: docker-build
docker-build:
	docker build -t hzhyvinskyi/iam-solutions-user-service:1.0.0 .

.PHONY: docker-run
docker-run:
	docker run -d -p 8089:8089 --name iam-solutions-user-service-container hzhyvinskyi/iam-solutions-user-service:1.0.0

.PHONY: docker-push
docker-push:
	docker push hzhyvinskyi/iam-solutions-user-service:1.0.0

.DEFAULT_GOAL := build-and-run
