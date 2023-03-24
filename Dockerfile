FROM golang:alpine AS builder
RUN apk update
RUN apk add musl-dev gcc
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w"

FROM alpine as runner
RUN apk update
RUN apk add sqlite
WORKDIR /

COPY --from=builder /game-server .

CMD ["./game-server"]
