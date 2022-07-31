# Dockerfile References:  https://docs.docker.com/engine/reference/builder/

FROM golang:1.18-alpine3.16

LABEL maintainer="Jonathan Throne <adramalech707@gmail.com>"

WORKDIR /go/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

EXPOSE 3000

CMD ./app
