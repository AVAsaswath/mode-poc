FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD . ./
RUN go build -o /docker-test

EXPOSE 8282
EXPOSE 3306


CMD [ "/docker-test" ]