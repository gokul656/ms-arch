build:
	./mvnw clean install -DskipTests=true
run:
	docker compose up --build -d
stop:
	docker compose down --rmi all