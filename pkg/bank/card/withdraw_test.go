package card

import (
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000

}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000

}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000

}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000

}