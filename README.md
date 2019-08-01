# fileserver

An alternative to `python -m http.server` in Go.

If you are unfamiliar with Python fear not, this project spins up a lightweight HTTP server to serve all static files in a folder.

## Requirements

- [Go](https://golang.org/) 1.12+
- [upx](https://upx.github.io/) (optional)

## How to build

```bash
go build -ldflags="-s -w" -o fileserver main.go
# Optional: Compress binary to ~75% of its original size (about 2MiB)
upx fileserver
```
