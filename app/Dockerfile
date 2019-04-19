FROM golang:1.11.4

WORKDIR /go/src/github.com/mamoroom/mock-go-clean-arch
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./

RUN go get github.com/pilu/fresh

RUN go mod download
RUN mkdir src
WORKDIR src

EXPOSE 80
