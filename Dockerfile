FROM golang:alpine AS builder
RUN apk update
RUN apk add musl-dev gcc
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM alpine as runner
RUN apk update
RUN apk add sqlite
WORKDIR /

COPY --from=builder /game-server .

ENV GIN_MODE=release
CMD ["./game-server"]
