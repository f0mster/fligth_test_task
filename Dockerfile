FROM golang:1.18-alpine as gotwirp
RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.2
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
RUN go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0

FROM pseudomuto/protoc-gen-doc:1.5.0 as protodoc

FROM golang:1.18-alpine

COPY --from=protodoc /usr/local/bin/protoc-gen-doc /usr/local/bin/protoc-gen-doc
COPY --from=gotwirp /go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=gotwirp /go/bin/protoc-gen-twirp /usr/local/bin/protoc-gen-twirp
COPY --from=gotwirp /go/bin/protoc-gen-go-vtproto /usr/local/bin/protoc-gen-go-vtproto

RUN apk add protoc

COPY ./gen.sh /usr/local/bin/gen.sh
ENTRYPOINT ["/usr/local/bin/gen.sh"]