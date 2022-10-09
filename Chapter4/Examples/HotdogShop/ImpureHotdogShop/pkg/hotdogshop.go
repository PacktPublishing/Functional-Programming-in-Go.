package pkg

const (
	HOTDOG_PRICE = 4
)

type CreditCard struct {
	credit int
}

type Hotdog struct{}

func (c *CreditCard) charge(amount int) {
	if amount <= c.credit {
		c.credit -= amount
	} else {
		panic("no more credit")
	}
}

func orderHotdog(c CreditCard) Hotdog {
	c.charge(HOTDOG_PRICE)
	return Hotdog{}
}
