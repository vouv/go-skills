FROM golang:1.14-alpine AS compile

MAINTAINER vouv

ENV HOST "127.0.0.1"

ENV PORT "3999"

EXPOSE 3999

WORKDIR /opt/go-skills

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

RUN go build -o application

ENTRYPOINT ["./application"]
