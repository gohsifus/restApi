FROM golang:1.16-alpine
RUN mkdir /restApi
ADD . /restApi
WORKDIR /restApi
RUN GOOS=linux GOARCH=arm64 go build -o main ./cmd/app
CMD ["/restApi/main"]