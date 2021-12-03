FROM golang:1.16-alpine as base

RUN mkdir /app
COPY . /app
WORKDIR /app

from base as test
CMD [ "go", "test" ]

from base as app
RUN go build -o main . 
CMD [ "/app/main" ]
