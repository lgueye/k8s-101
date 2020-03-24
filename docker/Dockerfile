FROM golang:alpine AS go-build
WORKDIR /build/app
COPY helloworld/go.mod helloworld/go.sum ./
RUN go mod download
COPY helloworld/main.go  ./
RUN go build -o helloworld .

FROM alpine
WORKDIR /app
COPY --from=go-build /build/app/helloworld /app/helloworld
RUN mkdir /logs
CMD ["./helloworld"]



