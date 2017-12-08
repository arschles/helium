.PHONY: docker-build-runtime
docker-build-runtime:
	GOOS=linux GOARCH=amd64 go build -o helium-runtime/runtime ./helium-runtime
	docker build -t helium-runtime:v0.0.1 helium-runtime

.PHONY: docker-build-runner
docker-build-runner:
	docker build -t helium-runner:v0.0.1 helium-runner
