package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return score/2 - 5
}

// Ability uses randomness to generate the score for an ability
func Ability() (sum int) {
	min := 6
	for i := 1; i <= 4; i++ {
		rnd := rand.Intn(6) + 1
		if min > rnd {
			min = rnd
		}
		sum += rnd
	}
	sum -= min
	return
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	con := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: con,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(con),
	}
}
