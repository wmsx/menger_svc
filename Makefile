.PHONY: docker
docker:
	docker build . -t menger-svc:latest
