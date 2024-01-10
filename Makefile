build:
	go build -o bin/app

run: build
	bin/app

db:
	docker run -d --name go-clean-postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Unclesnoopdog@69 -e POSTGRES_DB=go_clean -p 5432:5432 -v go_clean:/var/lib/postgresql/data postgres:latest
