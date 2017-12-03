FROM golang:1.9-alpine as builder

RUN mkdir -p /go/src/github.com/MOZGIII/discord-bot
WORKDIR /go/src/github.com/MOZGIII/discord-bot
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags '-s -extldflags "-static"' ./cmd/discord-bot

FROM alpine:3.6
RUN apk add --no-cache ffmpeg ca-certificates
COPY --from=builder /go/src/github.com/MOZGIII/discord-bot/discord-bot /usr/local/bin/discord-bot
CMD ["/usr/local/bin/discord-bot"]
