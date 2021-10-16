FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=offj
ADD ./src /build
RUN cd /build && go mod tidy && go build -o ./go-learning

CMD ["/build/go-learning"]