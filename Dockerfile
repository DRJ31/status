FROM golang:1.22-alpine as builder

ADD ./ /app

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
RUN GOARCH=amd64 GOOS=linux go build


FROM alpine

WORKDIR /app
COPY --from=builder /app/status /app
COPY public/ /app/public
RUN apk add --no-cache ca-certificates gcompat

CMD ["./status"]
EXPOSE 3000

