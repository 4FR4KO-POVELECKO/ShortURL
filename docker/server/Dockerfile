FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go server
RUN go mod download
RUN go build -o build -v ./cmd/server


CMD ["./build/server"]
