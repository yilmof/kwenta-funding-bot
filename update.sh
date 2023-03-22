#!/bin/bash

git pull

docker stop kwenta-funding-simulated

docker rmi kwenta_funding_simulated -f

docker build -t kwenta_funding_simulated -f Dockerfile .

docker rmi $(docker images -qa -f 'dangling=true')

docker rm kwenta-funding-simulated

docker run --name kwenta-funding-simulated -d --env-file ./.env kwenta_funding_simulated