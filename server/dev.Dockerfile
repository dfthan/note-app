FROM golang:1.19.3-alpine3.16

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

CMD air -c .air.toml
#CMD cd src && go run .


