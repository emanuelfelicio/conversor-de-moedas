package main

import (
	"fmt"
	"os"

	"github.com/emanuelfelicio/conversor-de-moedas/internal/ui"
)

func fatal(e error) {
	fmt.Fprintf(os.Stderr, "Erro: %v\n", e)
	os.Exit(1)
}
func main() {
	input, err := ui.RunForm()
	if err != nil {
		fatal(err)
	}
	err = input.Validate()
	if err != nil {
		fatal(err)
	}

}
