FROM golang:latest
ENV GOPATH=/

RUN go version

COPY ./ ./

RUN go mod download
RUN go build -v ./cmd/server

CMD ["./server"]
