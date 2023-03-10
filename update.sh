#!/bin/bash

git pull

docker stop kwenta-funding-bot

docker rmi kwenta_funding_bot -f

docker build -t kwenta_funding_bot -f Dockerfile .

docker rmi $(docker images -qa -f 'dangling=true')

docker rm kwenta-funding-bot

docker run --restart unless-stopped --name kwenta-funding-bot -d --env-file ./.env kwenta_funding_bot