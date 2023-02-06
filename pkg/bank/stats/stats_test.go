package stats

import (
	"fmt"

	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Category: "auto",
			Amount: 100_00,
		},
		{
			ID: 2,
			Amount: 300_00,
			Category: "dilivery",
		},
		{
			ID: 3,
			Amount: 300_00,
			Category: "hospital",
		},
		{
			ID: 4,
			Amount: 50_00,
			Category: "market",
		},
		
	}
	fmt.Println(Avg(payments))
	// Output:
	//  18750

}

func ExampleTotalinCategory() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Category: "auto",
			Amount: 100_00,
		},
		{
			ID: 2,
			Amount: 300_00,
			Category: "auto",
		},
		{
			ID: 3,
			Amount: 300_00,
			Category: "hospital",
		},
		{
			ID: 4,
			Amount: 50_00,
			Category: "market",
		},
		
	}
	fmt.Println(TotalinCategory(payments, "auto"))
	// Output:
	//  40000

}