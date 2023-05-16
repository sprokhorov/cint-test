.PHONY: test
test:
	go test -v ./...

.PHONY: docker-build
docker-build:
	docker buildx build --tag cint-test .
