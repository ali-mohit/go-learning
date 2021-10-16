FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=offj
ADD ./src /build
RUN cd /build && go mod init example.com/go_learning && go build --race .

CMD ["/build/go_learning"]