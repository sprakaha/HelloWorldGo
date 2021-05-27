FROM golang:1.16.0-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . .

RUN go build -o main . 

CMD ["/app/main"]