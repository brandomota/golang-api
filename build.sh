#! bin/sh

go build

if [ "$TRAVIS_BRANCH" != "master" ]; then
    docker build -t brandomota/golang-api-dev .
    docker push brandomota/golang-api-dev   
fi

if [ "$TRAVIS_BRANCH" = "master" ]; then
    docker build -t brandomota/golang-api-prod .
    docker push brandomota/golang-api-prod
if  