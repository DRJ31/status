FROM ubuntu:focal

WORKDIR /app
COPY status  /app
COPY public/ /app/public
RUN apt update
RUN apt install -y ca-certificates

CMD ["./status"]

EXPOSE 3000