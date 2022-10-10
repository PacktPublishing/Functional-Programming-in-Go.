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
			20,
			CreditCard{0},
			nil,
		},
		{
			CreditCard{150},
			1000,
			CreditCard{150},   // no money is withdrawn
			NOT_ENOUGH_CREDIT, // payment fails with this error
		},
	}
)

func TestCharge(t *testing.T) {
	for _, test := range testChargeStruct {
		t.Run("", func(t *testing.T) {
			output, err := Charge(test.inputCard, test.amount)
			if output != test.outputCard || !errors.Is(err, test.err) {
				t.Errorf("expected %v but got %v\n, error expected %v but got %v",
					test.outputCard, output, test.err, err)
			}
		})
	}
}

func TestOrderHotdog(t *testing.T) {
	testCC := CreditCard{1000}
	calledInnerFunction := false
	mockPayment := func(c CreditCard, input int) (CreditCard, CreditError) {
		calledInnerFunction = true
		testCC.credit -= input
		return testCC, nil
	}

	hotdog, resultF := OrderHotdog(testCC, mockPayment)
	if hotdog != NewHotdog() {
		t.Errorf("expected %v but got %v\n", NewHotdog(), hotdog)
	}

	_, err := resultF()
	if err != nil {
		t.Errorf("encountered %v but expected no error\n", err)
	}
	if calledInnerFunction == false {
		t.Errorf("Inner function did not get called\n")
	}
}
