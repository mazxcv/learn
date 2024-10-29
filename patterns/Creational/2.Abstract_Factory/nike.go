package main

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			"nike",
			15,
		},
	}
}

func (n *Nike) makShirt() IShirt {
	return &NikeShirt{
		Shirt{
			"nike",
			15,
		},
	}
}
