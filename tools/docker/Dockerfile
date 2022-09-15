FROM golang:1.18-alpine

RUN apk add protoc make

WORKDIR /src
ENV PWD=/src

COPY Makefile .

RUN make install-buf
RUN make install-go-deps

CMD ["make", "fast-generate"]
