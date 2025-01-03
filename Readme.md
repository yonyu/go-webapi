# Create Rest Web API in Golang

## Initial setup

        go mod init github.com/yonyu/go-webapi

        git remote add origin git@github.com:yonyu/go-webapi

        git checkout -b version-2-001
        git status
        git add -A
        git commit -m "Initial commit"
        (no longer supported) git push origin master

## Tools needed

go version

git version

docker -v 

docker-compose -v

task -h

## Architecture

                                                        Business Logic                       Database
                                                              ^
                                                              |
  Requests       >>>>               HTTP        -->        Service             -->           Repository
                                                              |
                                                              ⌵   
                                                         External APIs
                                                            Client

## Project layout

go-webapi/
   cmd/
       main.go
       server/

Open the terminal and run the command:         

go run cmd/main.go


go-webapi/
   internal/
        comment/
                comment.go

struct construction:
https://go.dev/doc/effective_go#composite_literals

Decouple the concrete implementation from database by returning struct

Add error handling

Add Implementation placeholder

Propagating context (Trace ID and Request ID), that is helpful when debugging multiple services

## Dockerizing the app

Create Dockerfile

        docker build -t go-webapi .
        docker run go-webapi

## Docker compose

Create docker-compose.yml file

        docker-compose up

## Create Taskfile.yml

Create Taskfile.yml file

        task build
        task test

## Connecting to database

        go get github.com/jmoiron/sqlx github.com/lib/pq

        task run

## Migrating the database

        go get github.com/golang-migrate/migrate/v4

to remove the pack by go get:

        go clean -i github.com/golang-migrate/migrate/v4

Step 1. create dbmigration.go and make other changes

Step 2. create migration SQL scripts

        task run

Open a new terminal:

        docker ps
        docker exec -it 95f8c33edaf4 bash
        psql -U postgres
        \dt
        \d+ comments;

## Implement database package

Step 1: Created comment.go

Step 2: 
        task run

another terminal:
        docker ps
        docker exec -it 95f8c33edaf4 bash
        psql -U postgres
        insert into comments (id) values('84f279ee-5aef-4ddb-8ae7-0d561a7944b2');
        select * from comments;
        \q
        exit
        
### Rename internal/comment/ to internal/domain/ 

Corresponding refactoring

### Rename internal/db/comment.go to internal/db/comment_repository.go

Rename internal/db/ to internal/database/

Rename internal/db/db.go to internal/database/database.go


Step 3: implement other database accesses

        go get github.com/satori/go.uuid

        task run

in a new terminal:
        docker exec -it 95f8c33edaf4 bash
        psql -U postgres
        select * from comments;
        \q
        exit

## Implement Delete and Updateoperations

DeleteComment

UpdateComment
