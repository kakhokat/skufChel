FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN apk update --no-cache && apk add --no-cache tzdata

RUN go build -o notifications ./cmd/main.go



FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow
WORKDIR /build

COPY --from=builder /build/notifications /build/notifications

COPY config/config.json /build/config/config.json
COPY migrations /build/migrations
CMD ["./notifications"]