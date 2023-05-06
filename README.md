# Relayroom

Simple game relay server written in go.

## Getting Started
```
version: "3"
services:
  game-server:
    image: sushiwaumai/relayroom
    ports:
      - 8080:8080
    env_file:
      .env
```

## License 
[MIT](./LICENSE)
