FROM golang:alpine AS builder
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM alpine as runner
WORKDIR /

COPY --from=builder /game-server .

CMD ["./game-server"]
