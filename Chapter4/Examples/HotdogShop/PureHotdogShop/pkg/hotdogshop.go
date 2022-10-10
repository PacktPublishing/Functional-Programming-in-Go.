package pkg

import "errors"

type CreditCard struct {
	credit int
}

type Hotdog struct {
	price int
}

type CreditError error

type PaymentFunc func(CreditCard, int) (CreditCard, CreditError)

var (
	NOT_ENOUGH_CREDIT CreditError = CreditError(errors.New("not enough credit"))
)

func NewCreditCard(initialCredit int) CreditCard {
	return CreditCard{credit: initialCredit}
}

func NewHotdog() Hotdog {
	return Hotdog{price: 4}
}

func Charge(c CreditCard, amount int) (CreditCard, CreditError) {
	if amount <= c.credit {
		c.credit -= amount
		return c, nil
	}
	return c, NOT_ENOUGH_CREDIT
}

func OrderHotdog(c CreditCard, pay PaymentFunc) (Hotdog, func() (CreditCard, error)) {
	hotdog := NewHotdog()
	chargeFunc := func() (CreditCard, error) {
		return pay(c, hotdog.price)
	}
	return hotdog, chargeFunc
}
