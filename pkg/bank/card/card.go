package card

import (
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
)


const limitRefill = 50_000_00
const MaxBonus = 5_000_00
func Deposit(card *types.Card, amount types.Money){
	if amount < 0 {
		return 
	}
	if amount > limitRefill {
		return 
	}
	if !(*card).Active {
		return 
	}
	
	(*card).Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	
	if !(*card).Active {
		return 
	}
	if (*card).Balance<0 {
		return
	}
	
	var Bonus types.Money = (*card).MinBalance*types.Money(percent)/100*types.Money(daysInMonth)/types.Money(daysInYear)

	if Bonus > types.Money(MaxBonus) {
		
		card.Balance += MaxBonus
		return 
	  }
	
	card.Balance += Bonus

	 
	
	
}

func Total(cards []types.Card) types.Money{
	sum:=types.Money(0)
	for _,	total := range cards{
		
		if total.Balance<0{
			continue
		}
		if !total.Active{
			continue
		}
	
		sum+=total.Balance
	}


	return sum
}

	
// Issue создает экземпляр карты с предопределенными полями
func Issue(currency types.Currency, color string, name string) types.Card	{
	return types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
}




