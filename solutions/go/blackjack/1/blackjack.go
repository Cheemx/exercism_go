package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
        "ace": 11,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "ten": 10,
        "jack": 10,
        "queen": 10,
        "king": 10,
        "other": 0,
    }
    return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
        case card1 == "ace" && card2 == "ace":
        	return "P"
        case ParseCard(card1) + ParseCard(card2) == 21 && !(ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11):
            return "W"
        case ParseCard(card1) + ParseCard(card2) == 21 && (ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11):
        	return "S"
        case ParseCard(card1) + ParseCard(card2) < 21 && ParseCard(card1) + ParseCard(card2) > 16:
        	return "S"
        case ParseCard(card1) + ParseCard(card2) > 11 && ParseCard(card1) + ParseCard(card2) < 17 && ParseCard(dealerCard) > 6:
        	return "H"
        case ParseCard(card1) + ParseCard(card2) < 12:
        	return "H"
        default:
        	return "S"
    }
    return "S"
}
