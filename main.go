package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chimajax/password-gen/generator"
)

func main() {
	length := flag.Int("length", 16, "password length")
	count := flag.Int("count", 1, "number of passwords to generate")
	noLower := flag.Bool("no-lower", false, "exclude lowercase letters")
	noUpper := flag.Bool("no-upper", false, "exclude uppercase letters")
	noNumbers := flag.Bool("no-numbers", false, "exclude numbers")
	noSymbols := flag.Bool("no-symbols", false, "exclude symbols")

	flag.Parse()

	opts := generator.Options{
		Length:  *length,
		Lower:   !*noLower,
		Upper:   !*noUpper,
		Numbers: !*noNumbers,
		Symbols: !*noSymbols,
	}

	for i := 0; i < *count; i++ {
		pw, err := generator.Generate(opts)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Println(pw)
	}
}
