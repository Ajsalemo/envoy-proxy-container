FROM golang:1.20.2-alpine3.17

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./application/go.mod ./
RUN go mod download && go mod verify

COPY ./application/ .
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["app"]