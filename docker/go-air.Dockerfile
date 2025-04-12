FROM golang:1.23-alpine

RUN mkdir app

ADD . /app/

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", "/app/.air.toml"]