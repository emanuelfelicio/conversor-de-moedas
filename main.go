package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emanuelfelicio/conversor-de-moedas/internal/domain"
	"github.com/emanuelfelicio/conversor-de-moedas/internal/exchange"
	"github.com/emanuelfelicio/conversor-de-moedas/internal/ui"
	"github.com/joho/godotenv"
)

const (
	EnvVarApiURL = "EXCHANGE_API_URL"
	EnvVarApiKey = "EXCHANGE_API_KEY"
	BaseApiURL   = "https://api.freecurrencyapi.com/v1/latest"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func fatal(e error) {
	fmt.Fprintf(os.Stderr, "Erro: %v\n", e)
	os.Exit(1)
}

func main() {
	input, err := ui.RunForm()
	if err != nil {
		fatal(err)
	}
	url := os.Getenv(EnvVarApiURL)
	if url == "" {
		url = BaseApiURL
	}
	key := os.Getenv(EnvVarApiKey)
	if key == "" {
		errorMsg := `Chave da API (EXCHANGE_API_KEY) não encontrada.
	1. Para obter uma chave gratuita, acesse: https://app.freecurrencyapi.com/register
	2. Crie um arquivo chamado '.env' na raiz do projeto.
	3. Adicione a seguinte linha ao arquivo .env:
	EXCHANGE_API_KEY="SUA_CHAVE_AQUI"
`
		fatal(fmt.Errorf("%s", errorMsg))
	}

	rate, err := exchange.Rate(url, key, input)
	if err != nil {
		fatal(err)
	}
	result, err := domain.Convert(input, rate)
	if err != nil {
		fatal(err)
	}
	ui.PrintResult(result)
}
