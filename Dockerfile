FROM golang:1.12.0-alpine3.9
WORKDIR /go/src/restApi
ADD . .
RUN apk add git
RUN go get -d -v
RUN go install -v
ENTRYPOINT /go/bin/restApi
EXPOSE 8080

