include compose/.env
include variables.mk

docker-image:
	make -C docker docker-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(service) ./cmd/

docker-start:
	make -C compose start

docker-stop:
	make -C compose stop

clean:
	@docker rmi -f $(service)
	@docker image prune -f
