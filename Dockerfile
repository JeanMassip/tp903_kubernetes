FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .
RUN go get -u
RUN go build main.go

EXPOSE 8083
ENTRYPOINT ["/app/main"]
