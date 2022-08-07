#!/bin/bash

# 1. start the postgres docker container

[ ! "$(docker ps -a | grep postgres)" ] && docker run --name postgres -d -p 5432:5432 \
    -e POSTGRES_USER=nft \
    -e POSTGRES_PASSWORD=nft \
    -e POSTGRES_DB=nft \
    library/postgres
echo "Postgres container started"

# 2. start the GraphQL server
go run main.go server.go


