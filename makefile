SERVICE_NAME := money-tracker-go

start:
	docker build -t $(SERVICE_NAME) . && \
	docker rm -f $(SERVICE_NAME) 2>/dev/null || true && \
	docker run --name $(SERVICE_NAME) -d -p 3000:3000 --env-file .env $(SERVICE_NAME)

