# Ignite 🔥

`ignite` is a simple and lightweight **Go hot-reload / live-reload tool & library** meant to help developers automatically rerun Go programs when source files change.

##  Usage

### CLI

```bash
go build -o ignite
./ignite "./path/to/**/*.go" "go run main.go"
