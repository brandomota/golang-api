#! /bin/bash

go build
docker build -t brandomota/golang-api-prod .
docker push brandomota/golang-api-prod
