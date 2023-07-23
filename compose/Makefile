DOCKER_COMPOSE=docker-compose
HAS_DOCKER_COMPOSE := $(shell command -v docker-compose 2> /dev/null)
ifndef HAS_DOCKER_COMPOSE
  DOCKER_COMPOSE=docker compose
  DOCKER_COMPOSE_QUIET_PULL=--quiet-pull
endif

include .env

stop:
	$(DOCKER_COMPOSE) -p$(docker) \
          -f docker-compose-dev.yml down


start:
	$(DOCKER_COMPOSE) -p$(docker) \
          -f docker-compose-dev.yml up $(DOCKER_COMPOSE_QUIET_PULL) -d

.PHONY: stop start
