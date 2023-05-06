FROM golang:alpine AS builder
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o a.out

FROM alpine as runner
WORKDIR /

COPY --from=builder /a.out .

ENV GIN_MODE=release
CMD ["./a.out"]
