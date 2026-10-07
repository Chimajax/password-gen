package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	symbolChars  = "!@#$%^&*()-_=+[]{};:,.<>?"
)

// Options controls what the generated password contains.
type Options struct {
	Length  int
	Lower   bool
	Upper   bool
	Numbers bool
	Symbols bool
}

// Generate creates a random password based on the given options.
func Generate(opts Options) (string, error) {
	if opts.Length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	var charset string
	if opts.Lower {
		charset += lowerChars
	}
	if opts.Upper {
		charset += upperChars
	}
	if opts.Numbers {
		charset += numberChars
	}
	if opts.Symbols {
		charset += symbolChars
	}

	if charset == "" {
		return "", errors.New("at least one character set must be enabled")
	}

	result := make([]byte, opts.Length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
