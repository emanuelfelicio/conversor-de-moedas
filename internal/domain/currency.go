package domain

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
)

var (
	ErrUnsupportedCurrency = errors.New("unsupported currency")
	ErrAmountIsNegative    = errors.New("amount must be >= 0")
	ErrInvalidRate         = errors.New("rate must be > 0")
)

// Amount is a Value Object representing a monetary value in cents.
type Amount int64

// NewAmountFromString creates an Amount from a string representation (e.g., "10.50").
func NewAmountFromString(s string) (Amount, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount format: %w", err)
	}
	cents := math.Round(f * 100)
	return Amount(cents), nil
}

// String formats the Amount for display (e.g., 1050 -> "10.50").
func (a Amount) String() string {
	f := float64(a) / 100.0
	return fmt.Sprintf("%.2f", f)
}

// Currency is a Value Object for currency codes.
type Currency string

var supported = []Currency{"BRL", "USD", "EUR"}

func (c Currency) isSupported() bool {
	return slices.Contains(supported, c)
}

func SupportedCurrency() []Currency {
	return supported
}

// ConversionInput encapsulates the input for a currency conversion.
type ConversionInput struct {
	from   Currency
	to     Currency
	amount Amount
}

// NewConversionInput creates a new ConversionInput
func NewConversionInput(from, to Currency, amount Amount) (*ConversionInput, error) {
	if !from.isSupported() {
		return nil, ErrUnsupportedCurrency
	}
	if !to.isSupported() {
		return nil, ErrUnsupportedCurrency
	}
	if amount < 0 {
		return nil, ErrAmountIsNegative
	}
	return &ConversionInput{from: from, to: to, amount: amount}, nil
}

// Getters for ConversionInput fields.
func (i *ConversionInput) From() Currency { return i.from }
func (i *ConversionInput) To() Currency   { return i.to }
func (i *ConversionInput) Amount() Amount { return i.amount }

// ConversionResult encapsulates the result of a currency conversion.
type ConversionResult struct {
	input  *ConversionInput
	rate   float64
	amount Amount
}

// Getters for ConversionResult fields.
func (r *ConversionResult) Input() *ConversionInput { return r.input }
func (r *ConversionResult) Rate() float64           { return r.rate }
func (r *ConversionResult) Amount() Amount          { return r.amount }

// Convert performs the currency conversion logic.
func Convert(input *ConversionInput, rate float64) (*ConversionResult, error) {
	if rate <= 0 {
		return nil, ErrInvalidRate
	}

	convertedAmountInCents := float64(input.Amount()) * rate
	roundedAmount := math.Round(convertedAmountInCents)
	converted := Amount(roundedAmount)

	return &ConversionResult{
		input:  input,
		rate:   rate,
		amount: converted,
	}, nil
}
