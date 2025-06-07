#!/bin/bash

echo 1. create configs from examples
cp ./config.json.example ./config.json
cp ./dbconfig.yml.example ./dbconfig.yml
cp ./docker-compose.env.example ./docker-compose.env

sleep 1s

echo 2. build and run
docker-compose --env-file docker-compose.env up --build -d

sleep 2s

echo 3. migrate
go install github.com/rubenv/sql-migrate/sql-migrate@latest

sql-migrate status
sql-migrate up

