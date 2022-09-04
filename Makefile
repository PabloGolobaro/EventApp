dev:
	docker-compose -f docker-compose.dev.yml up
dev_build:
	docker-compose -f docker-compose.dev.yml build

production:
	docker-compose -f docker-compose.yml up

prod_build:
	docker-compose -f docker-compose.yml build

#build:
#	docker build -t app-prod . --target production

#start:
#	docker run -p 8080:8080 --name app-prod app-prod