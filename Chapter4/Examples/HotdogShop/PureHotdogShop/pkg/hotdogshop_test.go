package pkg

import "testing"

var (
	testChargeStruct = []struct {
		inputCard  CreditCard
		amount     int
		outputCard CreditCard
		err        error
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
			CreditCard{},
			nil,
		},
	}
)

func testCharge(t *testing.T) {

	for _, test := range testChargeStruct {
		t.Run("", func(t *testing.T) {
			output, err := charge(test.inputCard, test.amount)
			if output != test.outputCard || (test.err == nil && err != nil) {
				t.Errorf("expected %v but got %v\n", test.outputCard, output)
			}
		})
	}

}
