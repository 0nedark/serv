.DEFAULT_GOAL := go

go: down up

down:
	docker-compose down

up:
	docker-compose up -d --remove-orphans

shell:
	docker exec -it serv-goconvey sh

logs:
	docker-compose logs -f --tail=100

.PHONY: go down up shell logs
