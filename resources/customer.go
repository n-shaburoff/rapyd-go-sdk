package resources

type Customer struct {
	Addresses            []Address     `json:"addresses"`
	BusinessVatID        string        `json:"business_vat_id"`
	Coupon               string        `json:"coupon"`
	DefaultPaymentMethod string        `json:"default_payment_method"`
	Description          string        `json:"description"`
	EWallet              string        `json:"e_wallet"`
	Name                 string        `json:"name"`
	Email                string        `json:"email"`
	Phone                string        `json:"phone_number"`
	PaymentMethod        PaymentMethod `json:"payment_method"`
}

type CustomerPaymentMethod struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type,omitempty"`
	Address  []Address              `json:"address"`
	Category string                 `json:"category"`
	Fields   map[string]interface{} `json:"fields"`
}

type CustomerResponse struct {
	Data Data `json:"data"`
}

type CustomerPaymentMethodListResponse struct {
	Data []CustomerPaymentMethod `json:"data"`
}
