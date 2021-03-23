FROM golang:1.16.2-alpine3.12 AS builder

WORKDIR /source
COPY . .

RUN mkdir bin
RUN GOOS=linux go build -o bin ./...

FROM alpine:3.12 AS service
ARG SERVICE_NAME

WORKDIR /app
COPY --from=builder /source/bin/$SERVICE_NAME app.bin

ENTRYPOINT ["./app.bin"]


