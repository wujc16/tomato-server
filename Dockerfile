FROM golang:1.17

ENV GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
ENV GO111MODULE=on

WORKDIR /server
COPY ./ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

ENV GIN_MODE=release

EXPOSE 8000
ENTRYPOINT ["./main"]
