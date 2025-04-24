# 🛠️ GoConfigParser

**GoConfigParser** is a lightweight command-line tool written in Go that parses configuration files written in an INI-like format. It extracts key-value pairs from structured config files and outputs them as Go structs or JSON.

---

## 📌 Problem Statement

> Build a robust, modular CLI tool in Go that reads configuration files, validates the syntax, and parses the data into structured Go types. The parser should support `[sections]` and key=value pairs, ignore comments and whitespace, and provide the output in JSON or Go-friendly formats.

This project is intended to practice and demonstrate:
- File I/O
- String manipulation
- CLI flag handling
- Basic parser logic
- Working with maps and structs
- Optional JSON output

---

## 📄 Input Format

Input config file (e.g., `config.cfg`):

```ini
# Server settings
[server]
host=127.0.0.1
port=8080

# Authentication credentials
[auth]
username=admin
password=secret
