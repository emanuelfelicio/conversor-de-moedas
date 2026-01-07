package ui

import (
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/emanuelfelicio/conversor-de-moedas/internal/domain"
)

func RunForm() (domain.ConversionInput, error) {

	var formData domain.ConversionInput
	var amount string
	currencyOptions := domain.SupportedCurrency()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[domain.Currency]().
				Title("Moeda de entrada").
				Options(huh.NewOptions(currencyOptions...)...).
				Value(&formData.From),

			huh.NewSelect[domain.Currency]().
				Title("Moeda de saída").
				Options(
					huh.NewOptions(currencyOptions...)...).
				Value(&formData.To),

			huh.NewInput().
				Title("Valor a converter").
				Placeholder("Ex: 100.50").Value(&amount),
		),
	)

	if err := form.Run(); err != nil {
		return domain.ConversionInput{}, err
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return domain.ConversionInput{}, err
	}

	formData.Amount = amountFloat
	return formData, nil
}
