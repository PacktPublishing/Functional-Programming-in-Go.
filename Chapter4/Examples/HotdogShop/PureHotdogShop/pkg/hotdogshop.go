package pkg

import "errors"

type CreditCard struct {
	credit int
}

type Hotdog struct {
	price int
}

func newHotdog() Hotdog {
	return Hotdog{price: 4}
}

func charge(c CreditCard, amount int) (CreditCard, error) {
	if amount <= c.credit {
		c.credit -= amount
		return c, nil
	}
	return c, errors.New("Could not charge credit card")
}

func orderHotdog(c CreditCard) (Hotdog, func() (CreditCard, error)) {
	hotdog := newHotdog()
	chargeFunc := func() (CreditCard, error) {
		return charge(c, hotdog.price)
	}
	return Hotdog{}, chargeFunc
}
