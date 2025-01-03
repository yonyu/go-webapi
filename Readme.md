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

## [Migrating the database](https://github.com/golang-migrate/migrate)

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


## Create the transport layer (containing the logic of transport)

Create internal/transport/http/handler.go

Scaffolding handler

        go get github.com/gorilla/mux

        task run

Launch a zsh terminal:

        curl http://localhost:8080/hello

### Gracefully shutdown

Run Server.ListenAndServe() in a Go routine.
The main routine is blocked until the channel receives a signal.

### Implement the handler function

Create transport/http/comment.go

        type Store interface {
                GetComment(context.Context, string) (Comment, error)
                PostComment(context.Context, Comment) (Comment, error)
                DeleteComment(context.Context, string) error
                UpdateComment(context.Context, string, Comment) (Comment, error)
        }

Implement PostComment()

in bash shell:

        task run

in zsh:

        curl --location --request POST 'http://localhost:8080/api/v1/comment' \
        --header 'Content-Type: application/json' \
        --data-raw '{
                "slug": "hello",
                "body": "body",
                "author": "me"
        }'

Implement GetComment()

        curl --location --request GET 'http://localhost:8080/api/v1/comment/2415b52e-3903-4c94-9bc4-4cef28fcf5aa'

Implment UpdateComment()

        curl --location --request PUT 'http://localhost:8080/api/v1/comment/2415b52e-3903-4c94-9bc4-4cef28fcf5aa' \
        --header 'Content-Type: application/json' \
        --data-raw '{ "Slug": "/testing-put", "Body": "body", "Author": "Yong" }'

Implement DeleteComment()

        curl --location --request DELETE 'http://localhost:8080/api/v1/comment/2415b52e-3903-4c94-9bc4-4cef28fcf5aa'


### Versioning Api endpoint

  For breaking changes, create a new version of the endpoint. For example:

        h.Router.HandleFunc("/api/v2/comment", h.PostCommentV2).Methods("POST")

### Implementing middleware

Create a file internal/transport/http/middleware.go

        curl --location --request POST 'http://localhost:8080/api/v1/comment' \
        --header 'Content-Type: application/json' -v\
        --data-raw '{
                "slug": "hello",
                "body": "body",
                "author": "me"
        }'


Handling error with middleware:

get request got the following error when using a non-existing comment ID:

        comments-rest-api  | Retrieve a domain
        comments-rest-api  | error fetching the domain by uuid: sql: no rows in result set
        comments-rest-api  | 2024/12/30 05:22:24 failed to fetch domain by ID


Logging middleware

        go get github.com/sirupsen/logrus

        to test:

        task run
        http://localhost:8080/api/v1/comment/2415b52e-3903-4c94-9bc4-4cef28fcf5aa

