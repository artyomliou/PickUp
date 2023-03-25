#!/bin/sh

if [ ! -e .env.docker ]; then
    sh ./setup.sh
fi

docker compose up -d
