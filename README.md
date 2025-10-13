# ignite 🔥

`ignite` is a simple and lightweight **Go hot-reload / live-reload tool & library** meant to help developers automatically rerun Go programs when source files change.

With `ignite` you can:
- Watch one or more file/glob path patterns for changes (create, write, remove)  
- Automatically restart the specified command (e.g. `go run main.go`)  
- Use as a CLI tool or integrate it as a library in your Go projects  

---

## 🚀 Features

- Minimal and easy to use  
- Hot restart of Go programs on code changes  
- Support glob / wildcard paths  
- Lightweight — no heavy dependencies  
- Can be extended (ignore rules, debounce, multiple paths, config files, etc.)

---

## 🛠️ Usage

### CLI

```bash
go build -o ignite
./ignite "./path/to/**/*.go" "go run main.go"
