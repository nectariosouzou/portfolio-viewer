build-dev:
	docker-compose -f ./docker-compose.dev.yml up

stop-dev:
	docker-compose -f ./docker-compose.dev.yml down

clean:
	make stop-dev
	docker image rm -f portfolio_viewer-frontend

build-prod:
	docker-compose -f ./docker-compose.prod.yml up --build -d
