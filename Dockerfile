FROM "scratch"

COPY golang-api golang-api

CMD [ "golang-api" ]