FROM golang:1.19.3-alpine3.16

WORKDIR /usr/src/app

COPY go.mod .
COPY . .

CMD cd src && go run .


