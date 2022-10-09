package pkg

import (
	"errors"
	"testing"
)

var (
	testChargeStruct = []struct {
		inputCard  CreditCard
		amount     int
		outputCard CreditCard
		err        CreditError
	}{
		{
			CreditCard{1000},
			500,
			CreditCard{500},
			nil,
		},
		{
			CreditCard{20},
			5,
			CreditCard{15},
			nil,
		},
		{
			CreditCard{20},
			1000,
			CreditCard{20},
			NOT_ENOUGH_CREDIT,
		},
	}
)

func TestCharge(t *testing.T) {
	for _, test := range testChargeStruct {
		t.Run("", func(t *testing.T) {
			output, err := charge(test.inputCard, test.amount)
			if output != test.outputCard || !errors.Is(err, test.err) {
				t.Errorf("expected %v but got %v\n, error expected %v but got %v", test.outputCard, output, test.err, err)
			}
		})
	}

}
