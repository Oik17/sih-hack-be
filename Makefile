run:
	go run cmd/main.go

postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_PASSWORD=Orchid7890 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres sih-hack

dropdb:
	docker exec -it postgres12 dropdb 
	
dockerfiledb:
	docker exec -it backend-db-1 bash

builddocker:
	docker build -t guptaakshat/caterpillar-hack-3:latest .

rundocker:
	docker run -p 8080:8080 guptaakshat/caterpillar-hack-3:latest

deploydocker:
	docker push guptaakshat/sih_hack:latest

tagdocker:
	docker tag sih_hack guptaakshat/sih_hack:latest

.PHONY: run postgres createdb dropdb dockerfiledb builddocker rundocker deploydocker