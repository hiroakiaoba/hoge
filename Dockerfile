FROM golang:1.12.4-alpine3.9

COPY . /go/src
WORKDIR /go/src

ENV GO111MODULE=on
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "/go/src/main.go"]