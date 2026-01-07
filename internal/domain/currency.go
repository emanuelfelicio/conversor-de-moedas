package domain

import "errors"

type Currency string

var supported = map[Currency]struct{}{
	"BRL": {},
	"USD": {},
	"EUR": {},
	"CNY": {},
	"JPY": {},
}

func (c Currency) isSupported() bool {
	_, ok := supported[c]
	return ok
}

func SupportedCurrency() []Currency {
	keys := make([]Currency, 0, len(supported))
	for k := range supported {
		keys = append(keys, k)
	}
	return keys
}

type ConversionInput struct {
	From   Currency
	To     Currency
	Amount float64
}

func (in ConversionInput) Validate() error {
	if !in.From.isSupported() {
		return errors.New("unsupported source currency")
	}
	if !in.To.isSupported() {
		return errors.New("unsupported target currency")
	}
	if in.Amount < 0 {
		return errors.New("amount must be >= 0")
	}
	return nil
}

type ConversionResult struct {
	From      Currency
	To        Currency
	Amount    float64
	Rate      float64
	Converted float64
}

func Convert(input ConversionInput, rate float64) (ConversionResult, error) {

	if err := input.Validate(); err != nil {
		return ConversionResult{}, err
	}

	if rate <= 0 {
		return ConversionResult{}, errors.New("rate must be > 0")
	}

	converted := input.Amount * rate

	return ConversionResult{
		From:      input.From,
		To:        input.To,
		Amount:    input.Amount,
		Rate:      rate,
		Converted: converted,
	}, nil
}
