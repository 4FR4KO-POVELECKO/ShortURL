FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go server
RUN go mod download
RUN make


CMD ["./build/server"]
