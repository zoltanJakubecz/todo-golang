FROM golang:1.17-bullseye

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["httpd"]
