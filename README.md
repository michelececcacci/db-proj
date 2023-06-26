DB Project

First, build the application image with `docker build -t go_application .`.

Then, run all the services with docker compose `docker compose up -d`.

After this step, just attach to the docker container named `go_application`
`docker attach *go_application*`.

Alternatively, just run 
```bash
export HOSTNAME=localhost 
export PORTNAME=55432
docker compose up -d
go run .
``` 
on a bash terminal

to populate the db, run 
```bash
go run faker/main/main.go
```

To reflect schema changes in the `db` image, first edit `schema.sql`, then run 
```bash
docker compose down --volumes
```
