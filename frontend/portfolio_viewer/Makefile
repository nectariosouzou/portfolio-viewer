build-dev:
	docker-compose -f ./docker-compose.dev.yml up

stop-dev:
	docker-compose down

clean:
	make stop-dev
	docker image rm -f frontend-dev