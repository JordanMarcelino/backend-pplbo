FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . /app

RUN go mod download && go mod tidy && go mod verify

RUN go build -v -o /app/binary ./cmd/main.go

FROM alpine:3.14

COPY --from=builder /app/binary /app/binary
COPY --from=builder /app/.env.prod /app.env

EXPOSE 8080
ENTRYPOINT /app/binary