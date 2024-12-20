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
        
