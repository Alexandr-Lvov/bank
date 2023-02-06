package payment
import(
	"fmt"
	"bank/pkg/bank/types"	
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