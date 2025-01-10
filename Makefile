include .env

BUILD_ARGS=--build-arg NEXT_PUBLIC_IMAGE_URL=$(NEXT_PUBLIC_IMAGE_URL)

build-migrate:
	@docker build -f ./api/Dockerfile.migrate -t grekas/jewellery-migrate api 

build-web:
	@docker build $(BUILD_ARGS) -t grekas/jewellery-web web

build-admin:
	@docker build $(BUILD_ARGS) -t grekas/jewellery-admin admin

build-api:
	@docker build -t grekas/jewellery-api api

push-web:
	@docker push grekas/jewellery-web

push-api:
	@docker push grekas/jewellery-api

push-admin:
	@docker push grekas/jewellery-admin

push-migrate:
	@docker push grekas/jewellery-migrate

build: build-web build-api build-admin build-migrate

push: push-web push-api push-admin push-migrate