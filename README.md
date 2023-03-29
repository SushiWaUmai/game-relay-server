# Game Relay Server

Simple game relay server written in go.

## Getting Started
```
version: "3"
services:
  game-server:
    image: sushiwaumai/game-relay-server
    ports:
      - 8080:8080
    env_file:
      .env
```

## License 
[MIT](./LICENSE)
