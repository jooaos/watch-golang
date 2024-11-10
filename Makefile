air-up:
	@air -c cmd/${service}/.air.toml

down:
	@docker compose down

up:
	@docker compose up --watch

up-no-cache:
	@docker compose up --build --watch

up-no-watch:
	@docker compose up -d

up-service:
	@docker compose up ${service} --watch -d 
