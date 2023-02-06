package card

import (
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
)

const limit = 20_000_00

func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount < 0 {
		return card
	}
	if amount > limit {
		return card
	}
	if card.Active == false {
		return card
	}
	if card.Balance < amount {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}
