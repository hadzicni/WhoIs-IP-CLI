# WhoIs IP CLI

A cross-platform command-line tool for querying WHOIS and IP geolocation data via public APIs.

## 🔧 Features

- Lookup WHOIS info for domain names
- Get geolocation data for IP addresses
- JSON output support
- Reverse DNS lookup for IPs
- CLI flags for version and help

## 📦 Build Instructions

```
go mod init whois-ip-cli
go build -o whoiscli.exe ./cmd/whoiscli
```

The executable `whoiscli.exe` will be created in the project root.

## 🚀 Usage

```
whoiscli <domain|ip> [flags]
```

### 🔹 Examples

```
whoiscli google.com
whoiscli 8.8.8.8
whoiscli -json google.com
whoiscli -json 8.8.8.8
whoiscli -v
whoiscli -h
```

## 🧾 Flags

| Flag    | Description           |
| ------- | --------------------- |
| `-json` | Output in JSON format |
| `-v`    | Show version info     |
| `-h`    | Show help message     |

## 🌍 Global Installation (optional)

To use `whoiscli` from anywhere:

1. Move `whoiscli.exe` to a directory like `C:\Tools\whoiscli\`
2. Add that directory to your Windows `PATH` environment variable

Then run:

```
whoiscli google.com
```

## 📁 Project Structure

```
whois-ip-cli/
├── cmd/
│   └── whoiscli/
│       └── main.go
├── internal/
│   └── whois/
│       ├── ip.go
│       └── domain.go
├── go.mod
├── README.md
└── LICENSE
```

## 📄 License

This project is licensed under the MIT License.
