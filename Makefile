BIN := blueprint

# Where to push the docker image.
REGISTRY ?= https://hub.docker.com/repositories

# This version-strategy uses git tags to set the version string
#VERSION := $(shell git describe --tags --always --dirty)

# This version-strategy uses a manual value to set the version string
VERSION := 1.0.10


dev:
	docker-compose -f docker-compose.dev.yml up
dev_build:
	docker-compose -f docker-compose.dev.yml build

production:
	docker-compose -f docker-compose.yml up

prod_build:
	docker-compose -f docker-compose.yml build

server:
	docker-compose up --no-start --force-recreate server
	docker push pablogolobar/notify_server:$(VERSION)
bot:
	docker-compose up --no-start --force-recreate bot
	docker push pablogolobar/notify_bot:$(VERSION)