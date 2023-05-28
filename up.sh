#!/bin/sh

sh ./setup.sh

docker compose up -d

cd backend
go run . -command migrate
go run . -command seed
go run . -command server
