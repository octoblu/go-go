FROM golang
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

RUN go get -u github.com/jteeuwen/go-bindata/...
ADD https://raw.githubusercontent.com/pote/gpm/v1.3.2/bin/gpm /go/bin/
RUN chmod +x /go/bin/gpm

COPY Godeps /go/src/github.com/octoblu/go-go/
WORKDIR /go/src/github.com/octoblu/go-go
RUN gpm install

COPY . /go/src/github.com/octoblu/go-go

RUN go-bindata -prefix templates templates/...
RUN env CGO_ENABLED=0 go build -a -ldflags '-s' .

CMD ["./go-go"]
