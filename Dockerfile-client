FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go client
RUN go mod download

RUN go build -o build -v ./cmd/client

CMD ["./build/client"]