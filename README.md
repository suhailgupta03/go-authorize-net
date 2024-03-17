# go-authorize-net

# Note:
This currently supports minimal authorize net methods and properties.

# Install
```shell
go get github.com/suhailgupta03/go-authorize-net
```

# Example
```go
package main

import (
	"fmt"
	"github.com/suhailgupta03/go-authorize-net/pkg/authorizenet"
)

func main() {
	transaction := authorizenet.New("apiLoginId", "transactionKey")
	transaction.AttachRefId("some-ref-911")

	tr := authorizenet.TransactionRequest{
		TransactionType: authorizenet.TransactionTypeAuthOnly,
		Amount:          9,
		CurrencyCode:    authorizenet.USD,
		Payment: &authorizenet.PaymentInformation{
			CreditCard: &authorizenet.CreditCard{
				CardNumber:     "5424000000000015",
				ExpirationDate: "2025-12",
				CardCode:       "999",
			},
		},
		PoNumber: "xyz-12345111",
		Customer: &authorizenet.Customer{
			ID:    "xyz189567",
			Email: "foo@bar",
		},
	}
	transaction.AttachTransactionRequest(tr)

	req := authorizenet.HTTPRequest{
		Test: true,
	}

	responseByte, err := req.AuthCreditCard(transaction)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(*responseByte))
		trans, err := authorizenet.ToTransactionResponse(responseByte)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Charging for ", trans.TransactionResponse.TransactionId)
			responseByte, err = req.CaptureAuthedAmount(transaction, trans.TransactionResponse.TransactionId)
			fmt.Println(string(*responseByte))
			f, _ := authorizenet.ToTransactionResponse(responseByte)
			fmt.Println(f.TransactionResponse.ResponseCode)

		}
	}
}
```