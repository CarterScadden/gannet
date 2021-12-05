FROM golang:1.16-alpine AS base
# used for go to be able to run the test, fails with an error about not finding cgo
# see: https://github.com/golang/go/issues/27303#issuecomment-504964623
ENV CGO_ENABLED 0
RUN mkdir /app
COPY . /app
WORKDIR /app

# the test image of the app can be built and ran with: `docker build . --target test -t test-image && docker run test-image`
FROM base AS test
RUN go test /app/services/produce
RUN go test /app/services
RUN go test /app/server/handlers

FROM base AS build
RUN go build -o main .

# the main image of the app, can be build and ran with `docker build . -t main-image && docker run -p <port to run on>:4000 main-image`
FROM build AS app
EXPOSE 4000
CMD [ "/app/main" ]
