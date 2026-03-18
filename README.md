# Network Lookup CLI

A command-line tool written in Go for querying network information — IP addresses, CNAME records, MX records, and name servers.

## Requirements

- Go 1.23.4+

## Installation

```bash
git clone https://github.com/windersegura/cli-tool
cd go-cli-sample
go build -o cli .
```

## Usage

```bash
./cli <command> --host <hostname>
```

The `--host` flag defaults to `google.com` if not provided.

## Commands

| Command | Description |
|---------|-------------|
| `ns`    | Look up name servers for a host |
| `ip`    | Look up IP addresses for a host |
| `cname` | Look up the CNAME for a host |
| `mx`    | Look up MX records for a host |

## Examples

```bash
# Name servers
./cli ns --host example.com

# IP addresses
./cli ip --host example.com

# CNAME
./cli cname --host www.example.com

# MX records
./cli mx --host example.com
```

## Dependencies

- [urfave/cli](https://github.com/urfave/cli) v1.22.17
