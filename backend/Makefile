build-backend:
	docker build -t redis-news-image .
	docker run -d -p 6379:6379 --name redis-news redis-news-image

run-backend:
	docker start redis-news
	go build
	./backend

stop-backend:
	docker stop redis-news
