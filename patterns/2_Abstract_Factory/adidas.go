package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			"adidas",
			14,
		},
	}
}

func (a *Adidas) makShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			"adidas",
			14,
		},
	}
}
