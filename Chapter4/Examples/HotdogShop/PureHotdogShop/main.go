package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter4/Examples/HotdogShop/PureHotdogShop/pkg"
)

func main() {
	myCard := pkg.NewCreditCard(1000)
	hotdog, creditFunc := pkg.OrderHotdog(myCard, pkg.Charge)
	fmt.Printf("%+v\n", hotdog)
	newCard, err := creditFunc()
	if err != nil {
		panic("User has no credit")
	}
	myCard = newCard
	fmt.Printf("%+v\n", myCard)
}
