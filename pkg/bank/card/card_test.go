package card
import (
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)

	fmt.Println(card.Balance)

	// Output: 3000000

}


func ExampleDeposit_noActive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)

	fmt.Println(card.Balance)

	// Output: 2000000

}

func ExampleDeposit_unlimit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)

	fmt.Println(card.Balance)

	// Output: 2000000

}

func ExampleAddBonus_positive() {
	card := types.Card{MinBalance: 20_000_00, Balance: 20_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2004931
}


func ExampleAddBonus_noActive() {
	card := types.Card{MinBalance: 20_000_00, Balance: 20_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_negative() {
	card := types.Card{MinBalance: 20_000_00, Balance: -2_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -200000
}

func ExampleAddBonus_maxBonus() {
	card := types.Card{MinBalance: 3000_000_00, Balance: 3000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 300500000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	
	}	

	fmt.Println(Total(cards))
	//Output: 2000000
	
}