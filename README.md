# h2cp

Proxy H2C to HTTP(s)

## Usage

```bash
$ h2cp -h

$ h2cp -addr=:8080 -target=https://example.com
$ h2cp -addr=:8080 -target=http://localhost:8080
$ h2cp -addr=:8080 -target=https://localhost:8443
$ h2cp -addr=:8080 -target=h2c://localhost:8080
$ h2cp -addr=:8080 -target=unix:///var/run/server.sock
```

## License

MIT
