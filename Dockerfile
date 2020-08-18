FROM "scratch"

COPY golang-api golang-api

EXPOSE 8080

CMD [ "golang-api" ]