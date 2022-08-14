dev:
	docker-compose -f docker-compose.dev.yml up

build:
	docker build -t app-prod . --target production

start:
	docker run -p 8080:8080 --name app-prod app-prod