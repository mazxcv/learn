package lasagna

func PreparationTime(layers []string, estimate int) int {
	if estimate == 0 {
		estimate = 2
	}
	return len(layers) * estimate
}

func Quantities(layers []string) (int, float64) {
	noodlesByLayers := 50
	noodles := "noodles"
	noodlesAmount := 0
	sauceByLayers := 0.2
	sauce := "sauce"
	sauceAmount := 0.0
	for _, v := range layers {
		if v == noodles {
			noodlesAmount += noodlesByLayers
		}
		if v == sauce {
			sauceAmount += sauceByLayers
		}
	}

	return noodlesAmount, sauceAmount
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, countPerson int) []float64 {
	scale := float64(countPerson) / 2.0
	scaleQuantities := []float64{}

	for _, v := range quantities {
		scaleQuantities = append(scaleQuantities, v*scale)
	}
	return scaleQuantities
}
