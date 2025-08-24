package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	res := 0
	switch card {
	case "ace":
		res = 11
	case "two":
		res = 2
	case "three":
		res = 3
	case "four":
		res = 4
	case "five":
		res = 5
	case "six":
		res = 6
	case "seven":
		res = 7
	case "eight":
		res = 8
	case "nine":
		res = 9
	case "ten", "jack", "queen", "king":
		res = 10
	}

	return res
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	hand := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)
	decision := ""

	switch {
	case card1 == "ace" && card1 == card2:
		decision = "P"
	case hand == 21 && ParseCard(dealerCard) < 10:
		decision = "W"
	case hand == 21 && ParseCard(dealerCard) >= 10:
		decision = "S"
	case hand >= 17 && hand <= 20:
		decision = "S"
	case hand >= 12 && hand <= 16 && dealerHand < 7:
		decision = "S"
	case hand >= 12 && hand <= 16 && dealerHand >= 7:
		decision = "H"
	case hand <= 11:
		decision = "H"
	}

	return decision

}
