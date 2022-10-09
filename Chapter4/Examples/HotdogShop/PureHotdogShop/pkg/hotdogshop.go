package pkg

import "errors"

type CreditCard struct {
	credit int
}

type Hotdog struct {
	price int
}

type CreditError error

var (
	NOT_ENOUGH_CREDIT CreditError = CreditError(errors.New("not enough credit"))
)

func newHotdog() Hotdog {
	return Hotdog{price: 4}
}

func charge(c CreditCard, amount int) (CreditCard, CreditError) {
	if amount <= c.credit {
		c.credit -= amount
		return c, nil
	}
	return c, NOT_ENOUGH_CREDIT
}

func orderHotdog(c CreditCard) (Hotdog, func() (CreditCard, error)) {
	hotdog := newHotdog()
	chargeFunc := func() (CreditCard, error) {
		return charge(c, hotdog.price)
	}
	return Hotdog{}, chargeFunc
}
