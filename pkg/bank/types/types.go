package types

// Money представляет денежную сумму в минимальных единицах(центы, копейки, дирамы и т.д.)
type Money int64
// Category  представляет собой категорию, в которой был совершен платеж(авто, аптеки, рестораны, и т.д.)
type Category string


type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
	MinBalance Money
}

// Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
	Category Category
}

// PaymentSource предствляет источник оплаты
type PaymentSource struct {
	Type    string // "card"
	Number  string // номер вида "5058 xxxx xxxx 8888"
	Balance Money  // баланс в дирамах
}
