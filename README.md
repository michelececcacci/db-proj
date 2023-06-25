DB Project

First, build the application image with `docker build -t go_application .`.

Then, run all the services with docker compose `docker compose up -d`.

After this step, just attach to the docker container named `go_application`
`docker attach *go_application*`.

Alternatively, just run 
```bash
HOSTNAME=localhost 
PORTNAME=55432
go run .
``` 
on a bash terminal

to populate the db, run 
```go run  faker/main/main.go
```
