FROM golang:1.20-alpine3.17 AS dev

WORKDIR /app/tcp_server

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY server ./server

RUN go build -o app ./server/main.go

ENV SERVER_PORT=${SERVER_PORT}
EXPOSE 4040
CMD [ "go" "run" "./server/main"]

FROM alpine:3.17.2 AS prod
WORKDIR /app/tcp_server

COPY --from=dev /app/tcp_server/app ./

CMD [ "./app" ]
