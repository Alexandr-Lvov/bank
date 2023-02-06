package payment

import (
	"github.com/Alexandr-Lvov/bank/pkg/bank/types"
)

func Max(payments []types.Payment)types.Payment{
	max:=types.Payment(payments[0])
	for _,	payment := range payments{
	
		if max.Amount<payment.Amount {
			max=payment
		}
	}
	
	return max
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	result := make([]types.PaymentSource, 0, len(cards))

	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

		
		paySource := types.PaymentSource{
			Type:    card.Name,
			Number:  string(card.PAN),
			Balance: card.Balance,
		}
		result = append(result, paySource)
	}

	return result
}