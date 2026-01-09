package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/emanuelfelicio/conversor-de-moedas/internal/domain"
)

func RunForm() (*domain.ConversionInput, error) {
	var fromCurrency domain.Currency
	var toCurrency domain.Currency
	var amountStr string
	currencyOptions := domain.SupportedCurrency()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[domain.Currency]().
				Title("Moeda de entrada").
				Options(huh.NewOptions(currencyOptions...)...).
				Value(&fromCurrency),

			huh.NewSelect[domain.Currency]().
				Title("Moeda de saída").
				Options(
					huh.NewOptions(currencyOptions...)...).
				Value(&toCurrency),

			huh.NewInput().
				Title("Valor a converter").Prompt("?").
				Placeholder("Ex: 100.50").Value(&amountStr),
		),
	).WithTheme(huh.ThemeDracula())

	if err := form.Run(); err != nil {
		return nil, err
	}

	amount, err := domain.NewAmountFromString(amountStr)
	if err != nil {
		return nil, err
	}

	input, err := domain.NewConversionInput(fromCurrency, toCurrency, amount)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func PrintResult(r *domain.ConversionResult) {
	fmt.Printf("%s %s -> %s %s\n", r.Input().Amount().String(), r.Input().From(), r.Amount().String(), r.Input().To())
}
