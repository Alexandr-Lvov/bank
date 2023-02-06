package payment
import (
	"fmt"
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"	
)
func ExampleMax() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Amount: 100_00,
		},
		{
			ID: 2,
			Amount: 300_00,
		},
		{
			ID: 3,
			Amount: 300_00,
		},
		{
			ID: 4,
			Amount: 50_00,
		},
		
	}

	max:=Max(payments)
	fmt.Println(max.ID)
	//Output: 2
	
}

func ExamplePaymentSources_positive() {
	cards := []types.Card{
		{PAN: "1111 xxxx xxxx 9999",
		 Balance: 44,
		 Active: true,
		},
		{PAN: "2222 xxxx xxxx 8888",
		Balance: 47411,
		Active: true,
		},
	}

	result := PaymentSources(cards)
	for _, paySource := range result {
		fmt.Println(paySource.Number)
	}

	// Output:
	// 1111 xxxx xxxx 9999
	// 2222 xxxx xxxx 8888
}

func ExamplePaymentSources_noMoney() {
	cards := []types.Card{
		{PAN: "1111 xxxx xxxx 9999",
		 Balance: 0,
		 Active: true,
		},
		{PAN: "2222 xxxx xxxx 8888",
		Balance: 47411,
		Active: true,
		},
	}

	result := PaymentSources(cards)
	for _, paySource := range result {
		fmt.Println(paySource.Number)
	}

	// Output:
	// 2222 xxxx xxxx 8888
}
func ExamplePaymentSources_inactive() {
	cards := []types.Card{
		{PAN: "1111 xxxx xxxx 9999",
		 Balance: 44,
		 Active: true,
		},
		{PAN: "2222 xxxx xxxx 8888",
		Balance: 47411,
		Active: false,
		},
	}

	result := PaymentSources(cards)
	for _, paySource := range result {
		fmt.Println(paySource.Number)
	}

	// Output:
	// 1111 xxxx xxxx 9999
	}