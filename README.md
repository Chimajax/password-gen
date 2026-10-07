# Password-gen

[![Test](https://github.com/Chimajax/password-gen/actions/workflows/test.yml/badge.svg)](https://github.com/Chimajax/password-gen/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Chimajax/password-gen)](https://goreportcard.com/report/github.com/Chimajax/password-gen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A small, cryptographically secure password generator CLI written in Go.
No dependencies — just the standard library.

## Features

- 🔐 Uses `crypto/rand` — not `math/rand`. Real cryptographic randomness.
- ⚙️ Configurable length, character sets, and count
- 📦 Zero external dependencies
- ✅ Fully testdd
- 🚀 Tiny binary (builds in under a second)

## Install

```bash
go install github.com/Chimajax/password-gen@latest
```
or from source

```bash
git clone https://github.com/Chimajax/password-gen.git
cd password-gen
make build
```

## Usage

```bash

    # Generate one 16-character password (default)
password-gen

# Generate a 24-character password
password-gen -length 24

# Generate 5 passwords at once
password-gen -length 32 -count 5

# No symbols (for sites that don't allow them)
password-gen -no-symbols

# Digits only
password-gen -no-lower -no-upper -no-symbols

# Alphanumeric only
password-gen -no-symbols

```

### Example Output
```bash
$ password-gen -length 20 -count 3
fE1=P]2S:h_Ms5&kZ<yP
#W9Db+VvA_3xmauh>8])
JX%=Ad0JA<x+D;lyM4HL
```

## Why crypto/rand?
Passwords generated with math/rand are predictable — given enough output, an
attacker can reconstruct the seed and guess future passwords. This tool uses
crypto/rand, which reads from the operating system's cryptographic entropy
source. It's the same source used by TLS, SSH, and your OS keychain.

## Flags

| Flag         | Default | Description                   |
|--------------|---------|-------------------------------|
| -length      | 16      | password length               |
| -count       | 1       | how many to generate          |
| -no-lower    | false   | exclude lowercase             |
| -no-upper    | false   | exclude uppercase             |
| -no-numbers  | false   | exclude numbers               |
| -no-symbols  | false   | exclude symbols               |

## Tests

    make test

## License

MIT — see [LICENSE](license/)
