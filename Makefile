SHELL=/bin/bash
SRC=./main.go ./util/ ./view/ ./queries
$DIRMAKE=./.makefile/

./queries: ./query.sql ./schema.sql ./sqlc.yaml ./.makefile/d-volumes
	sqlc generate

.PHONY: d-volumes
d-volumes: ./.makefile/d-volumes

./.makefile/d-volumes: 
	sudo docker compose down --volumes
	touch ./.makefile/d-volumes

.PHONY: d-build
d-build: ./.makefile/d-build

./.makefile/d-build: ${SRC} ./go.mod ./go.sum ./Dockerfile 
	sudo docker build -t go_application .
	touch ./.makefile/d-build

.PHONY: up
up: ./.makefile/d-build
	sudo docker compose up -d

.PHONY: down 
down:
	sudo docker compose down
	
.PHONY: attach
attach: up 
	sudo docker attach db-proj-go_application-1
 
