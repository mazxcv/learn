package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}

}
