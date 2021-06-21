FROM golang:buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o server ./cmd/server/main.go

CMD ["./server"]
