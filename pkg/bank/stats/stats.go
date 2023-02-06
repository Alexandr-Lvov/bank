package stats

import (
	

	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
)

// Avg рассчитыввает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var total types.Money
	
	for _, payment := range payments {
	total += payment.Amount
	
	}

	return total/types.Money(len(payments))
}

// TotalinCategory находит сумму покупок в определенной категории
func TotalinCategory(payments []types.Payment, category types.Category) types.Money{
	var totalCategory types.Money
	for _, payment:= range payments {
		if payment.Category == category{
			totalCategory +=payment.Amount
		}
		
	}
	return totalCategory
}