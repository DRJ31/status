FROM ubuntu

WORKDIR /app
COPY tiebarankgo /app
RUN mkdir /app/public
COPY index.html /app/public
RUN apt update
RUN apt install -y ca-certificates

CMD ["./status"]

EXPOSE 80