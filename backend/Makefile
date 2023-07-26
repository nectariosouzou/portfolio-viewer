run-app:
	docker-compose up --force-recreate --build server

build-backend:
	docker-compose up --force-recreate --build server

stop-backend:
	docker-compose down || true

clean:
	make stop-backend
	docker image rm -f redis:alpine
	docker image rm -f backend-server


