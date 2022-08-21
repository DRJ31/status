FROM alpine

WORKDIR /app
COPY status  /app
COPY public/ /app/public
RUN apk add --no-cache gcompat ca-certificates

CMD ["./status"]

EXPOSE 3000
