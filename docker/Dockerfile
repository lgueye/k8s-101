FROM golang:latest
WORKDIR /app
COPY helloworld/go.mod helloworld/go.sum ./
RUN go mod download
COPY helloworld/main.go .
RUN go build -o helloworld .
CMD ["./helloworld"]



