FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/{{.ProjectName}}
COPY . /go/src/github.com/octoblu/{{.ProjectName}}

RUN env CGO_ENABLED=0 go build -o {{.ProjectName}} -a -ldflags '-s' .

CMD ["./{{.ProjectName}}"]
