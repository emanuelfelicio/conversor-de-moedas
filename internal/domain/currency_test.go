package domain

import (
	"errors"
	"testing"
)

func TestNewConversionInput(t *testing.T) {
	tests := []struct {
		name          string
		from, to      Currency
		amount        Amount
		expectedErr   error
		expectedInput *ConversionInput
	}{
		{
			name:          "Invalid source currency",
			from:          "JPY",
			to:            "BRL",
			amount:        30,
			expectedErr:   ErrUnsupportedCurrency,
			expectedInput: nil,
		},
		{
			name:          "Invalid target currency",
			from:          "USD",
			to:            "ABC",
			amount:        500,
			expectedErr:   ErrUnsupportedCurrency,
			expectedInput: nil,
		},
		{
			name:          "Negative amount",
			from:          "USD",
			to:            "BRL",
			amount:        -100,
			expectedErr:   ErrAmountIsNegative,
			expectedInput: nil,
		},
		{
			name:        "Valid input",
			from:        "USD",
			to:          "BRL",
			amount:      Amount(1050),
			expectedErr: nil,
			expectedInput: &ConversionInput{
				from:   "USD",
				to:     "BRL",
				amount: Amount(1050),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := NewConversionInput(tt.from, tt.to, tt.amount)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("NewConversionInput() got error = %v, wantErr %v", err, tt.expectedErr)
			}

			if tt.expectedErr == nil {
				if input == nil {
					t.Errorf("NewConversionInput() got nil input, want non-nil")
				} else {
					if input.From() != tt.expectedInput.From() {
						t.Errorf("NewConversionInput() From() = %v, want %v", input.From(), tt.expectedInput.From())
					}
					if input.To() != tt.expectedInput.To() {
						t.Errorf("NewConversionInput() To() = %v, want %v", input.To(), tt.expectedInput.To())
					}
					if input.Amount() != tt.expectedInput.Amount() {
						t.Errorf("NewConversionInput() Amount() = %v, want %v", input.Amount(), tt.expectedInput.Amount())
					}
				}
			} else {
				if input != nil {
					t.Errorf("NewConversionInput() got input %v, want nil", input)
				}
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		input       *ConversionInput
		rate        float64
		expectedErr error
		expectedAmt Amount
	}{
		{
			name:        "Happy path conversion and rounding",
			input:       mustNewConversionInput(t, "USD", "BRL", Amount(1050)), // 10.50 USD
			rate:        5.4321,
			expectedErr: nil,
			expectedAmt: Amount(5704), // 10.50 * 5.4321 = 57.03705 -> 57.04 (5704 cents)
		},
		{
			name:        "Exact conversion",
			input:       mustNewConversionInput(t, "USD", "USD", Amount(1000)), // 10.00 USD
			rate:        1.00,
			expectedErr: nil,
			expectedAmt: Amount(1000),
		},
		{
			name:        "Rounding half up",
			input:       mustNewConversionInput(t, "EUR", "BRL", Amount(100)), // 1.00 EUR
			rate:        1.50,
			expectedErr: nil,
			expectedAmt: Amount(150),
		},
		{
			name:        "Rounding below half",
			input:       mustNewConversionInput(t, "EUR", "BRL", Amount(100)), // 1.00 EUR
			rate:        1.49,
			expectedErr: nil,
			expectedAmt: Amount(149),
		},
		{
			name:        "Rounding above half",
			input:       mustNewConversionInput(t, "EUR", "BRL", Amount(100)), // 1.00 EUR
			rate:        1.51,
			expectedErr: nil,
			expectedAmt: Amount(151),
		},
		{
			name:        "Invalid rate (zero)",
			input:       mustNewConversionInput(t, "USD", "BRL", Amount(100)),
			rate:        0,
			expectedErr: ErrInvalidRate,
			expectedAmt: 0,
		},
		{
			name:        "Invalid rate (negative)",
			input:       mustNewConversionInput(t, "USD", "BRL", Amount(100)),
			rate:        -1.0,
			expectedErr: ErrInvalidRate,
			expectedAmt: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.input, tt.rate)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Convert() got error = %v, wantErr %v", err, tt.expectedErr)
			}

			if tt.expectedErr == nil {
				if result == nil {
					t.Errorf("Convert() got nil result, want non-nil")
				} else {
					if result.Amount() != tt.expectedAmt {
						t.Errorf("Convert() result Amount() = %v, want %v", result.Amount(), tt.expectedAmt)
					}
					// Also check input is correctly stored
					if result.Input() != tt.input {
						t.Errorf("Convert() result Input() = %v, want %v", result.Input(), tt.input)
					}
				}
			} else {
				if result != nil {
					t.Errorf("Convert() got result %v, want nil", result)
				}
			}
		})
	}
}

// Helper function to create ConversionInput in tests
func mustNewConversionInput(t *testing.T, from, to Currency, amount Amount) *ConversionInput {
	input, err := NewConversionInput(from, to, amount)
	if err != nil {
		t.Fatalf("Failed to create ConversionInput: %v", err)
	}
	return input
}
