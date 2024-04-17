package authorizenet

type MerchantAuthentication struct {
	// Merchant’s unique API Login ID.
	Name string `json:"name"`
	// Merchant’s unique Transaction Key.
	TransactionKey string `json:"transactionKey"`
}

type CreditCard struct {
	CardNumber string `json:"cardNumber,omitempty"`
	// ExpirationDate (YYYY-MM) formatting.
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type PaymentInformation struct {
	CreditCard *CreditCard `json:"creditCard,omitempty"`
}

type TransactionRequest struct {
	TransactionType string              `json:"transactionType"`
	Amount          float64             `json:"amount"`
	CurrencyCode    string              `json:"currencyCode"`
	RefTransId      string              `json:"refTransId,omitempty"`
	Payment         *PaymentInformation `json:"payment,omitempty"`
	// The merchant-assigned purchase order number.
	PoNumber string    `json:"poNumber,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	BillTo   *BillTo   `json:"billTo,omitempty"`
}

type Customer struct {
	// The unique customer ID used to represent the customer associated with the transaction.
	// String, up to 20 characters.
	ID string `json:"id,omitempty"`
	// The customer’s valid email address. String, up to 255 characters.
	Email string `json:"email,omitempty"`
}

type BillTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type Request struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	// Merchant-assigned reference ID for the request.
	RefId              string             `json:"refId"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

// CustomerProfile used with CustomerProfileRequest
type CustomerProfile struct {
	MerchantCustomerId string `json:"merchantCustomerId"`
	Description        string `json:"description"`
	Email              string `json:"email"`
}
type CreateCustomerProfileRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	Profile                CustomerProfile        `json:"profile"`
	ValidationMode         string                 `json:"validationMode,omitempty"`
}

type CreateCustomerProfile struct {
	CreateCustomerProfileRequest CreateCustomerProfileRequest `json:"createCustomerProfileRequest"`
}

type Transaction struct {
	CreateTransactionRequest Request `json:"createTransactionRequest"`
}
