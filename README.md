# 🌐 WhoIs IP CLI

A fast and minimal CLI tool written in Go to fetch **Whois information for domains** and **IP geolocation details** directly from the terminal. Supports JSON output and versioning. Powered by public APIs.

![Go Version](https://img.shields.io/badge/Go-1.24+-blue?logo=go)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🌍 Lookup domain Whois info via `api.whois.vu`
- 📍 Get IP geolocation and provider info via `ip-api.com`
- 🔎 Detects input type (IP or domain) automatically
- 📦 JSON output option for automation or scripting
- 🧾 Simple flags: `-json`, `-v`, `-h`
- ⚙️ Written in pure Go with zero dependencies

---

## 📦 Installation

### Option 1: Go Install

```bash
go install github.com/hadzicni/whois-ip-cli/cmd/whoiscli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build

```bash
git clone https://github.com/hadzicni/whois-ip-cli.git
cd whois-ip-cli/cmd/whoiscli
go build -o whoiscli
```

---
## 🚀 Usage

```bash
whoiscli [flags] <domain|ip>
```

### Available Flags

| Flag        | Description                 |
|-------------|-----------------------------|
| `-json`     | Output as JSON              |
| `-v`        | Show version info           |
| `-h`        | Show help message           |

---

## 🔧 Examples

Check a domain:

```bash
whoiscli example.com
```

Check an IP address:

```bash
whoiscli 8.8.8.8
```

Output as JSON:

```bash
whoiscli -json example.com
```

Show version:

```bash
whoiscli -v
```

---

## 👨‍💻 Author

Made by **Nikola Hadzic**  
GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
