# 🌐 WhoIs IP CLI

A fast and simple command-line tool to perform WHOIS lookups for IP addresses and domains — directly from your terminal.

![Go Version](https://img.shields.io/badge/Go-1.20+-blue?logo=go)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🔎 Lookup WHOIS information for IP addresses or domains  
- ⚡ Super fast response using Go’s built-in networking libraries  
- 📦 Lightweight CLI with zero external dependencies  
- 🧠 Intelligent input detection (domain or IP)  
- 🧪 Clean output format for readability  

---

## 📦 Installation

### Option 1: Go Install (recommended)

```bash
go install github.com/hadzicni/whois-ip-cli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build

```bash
git clone https://github.com/hadzicni/whois-ip-cli.git
cd whois-ip-cli
go build -o whois-ip ./cmd/whois-ip
```

---

## 🚀 Usage

```bash
whois-ip [flags] <target>
```

### Parameters

| Parameter        | Description                          | Example                     |
|------------------|--------------------------------------|-----------------------------|
| `target`         | IP address or domain to look up      | `whois-ip 8.8.8.8`          |

### Flags

| Flag             | Description                          | Example                     |
|------------------|--------------------------------------|-----------------------------|
| `--help`, `-h`   | Show help message                    | `whois-ip -h`               |

---

## 🔧 Examples

WHOIS lookup for an IP address:

```bash
whois-ip 1.1.1.1
```

WHOIS lookup for a domain:

```bash
whois-ip example.com
```

---

## 🧪 Development

Run the CLI locally during development:

```bash
go run ./cmd/whois-ip example.com
```

Run tests:

```bash
go test ./...
```

---

## 📁 Project Structure

```
.
├── cmd/
│   └── whois-ip/        # CLI command logic
├── internal/
│   └── resolver/        # WHOIS resolution logic
├── go.mod               # Go module definition
├── LICENSE              # License file
└── README.md            # Project documentation
```

---

## 👨‍💻 Author

Made with ❤️ by **Nikola Hadzic**

- GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
