package yoocommon

type Item struct {
	// parameter with the mode of the payment
	PaymentMode string `json:"payment_mode"`
	
	// parameter with the type of the payment
	PaymentSubject string `json:"payment_subject"`
	
	// parameter with the name of the product or service
	Description string `json:"description"`

	// parameter with the amount per unit of product
	Quantity string `json:"quantity"`

	// parameter specifying the quantity of goods (only integers, for example 1)
	Amount *Amount `json:"amount"`

	// parameter with the fixed value 1 (price without VAT)
	VatCode string `json:"vat_code"`
}
