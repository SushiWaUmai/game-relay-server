FROM golang:alpine AS builder
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM alpine as runner
WORKDIR /

COPY --from=builder /game-relay-server .

ENV GIN_MODE=release
CMD ["./game-relay-server"]
