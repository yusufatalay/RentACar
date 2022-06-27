#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o yolcuapp .

EXPOSE 3000
CMD ["/app/yolcuapp"]
