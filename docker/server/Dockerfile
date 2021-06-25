FROM golang:latest

RUN go version
ENV GOPATH=/

COPY go.mod ./
RUN go mod download

COPY ./ ./
RUN go build -v ./cmd/server


CMD ["./server"]
