FROM golang:1.19.0

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 5000

CMD [ "/app/main" ]
