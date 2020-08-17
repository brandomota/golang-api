FROM "scratch"

COPY golang-api golang-api

EXPOSE 80

CMD [ "golang-api" ]