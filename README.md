# password-gen

A simple, cryptographically secure password generator written in Go.

## Install

    go install github.com/YOUR_USERNAME/password-gen@latest

## Usage

    password-gen -length 24 -count 3

    # exclude symbols
    password-gen -no-symbols

    # only numbers
    password-gen -no-lower -no-upper -no-symbols

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

MIT
