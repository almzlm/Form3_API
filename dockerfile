FROM golang:1.7-alpine

COPY . /go/src
WORKDIR /go/src/clientapp/clients_service

CMD ["go","test","-v"]