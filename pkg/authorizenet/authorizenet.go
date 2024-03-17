package authorizenet

import (
	"errors"
	"strings"
)

const (
	TransactionTypeAuthCapture      = "authCaptureTransaction"
	TransactionTypeAuthOnly         = "authOnlyTransaction"
	TransactionTypePriorAuthCapture = "priorAuthCaptureTransaction"
	USD                             = "usd"
)

const (
	incorrectPOString = "purchase order number must be 1-25 chars long"
	invalidCustomerId = "customer id must be 1-20 chars long"
)

func New(apiLoginId, transactionKey string) *Transaction {
	var transaction = Transaction{
		CreateTransactionRequest: Request{
			MerchantAuthentication: MerchantAuthentication{
				Name:           apiLoginId,
				TransactionKey: transactionKey,
			},
		},
	}

	return &transaction
}

func (tr *Transaction) AttachRefId(refId string) error {
	tr.CreateTransactionRequest.RefId = refId
	return nil
}

func (tr *Transaction) AttachTransactionRequest(req TransactionRequest) error {
	tr.CreateTransactionRequest.TransactionRequest = req
	return nil
}

func (tr *Transaction) AttachPONumber(poString string) error {
	po := strings.TrimSpace(poString)
	if po == "" || len(po) > 25 {
		return errors.New(incorrectPOString)
	}

	tr.CreateTransactionRequest.TransactionRequest.PoNumber = poString
	return nil
}

func (tr *Transaction) AttachCustomerDetails(customer *Customer) error {
	cid := strings.TrimSpace(customer.ID)
	if len(cid) == 0 || len(cid) > 20 {
		return errors.New(invalidCustomerId)
	}

	tr.CreateTransactionRequest.TransactionRequest.Customer = customer
	return nil
}

func (tr *Transaction) AttachTransactionType(transactionType string) error {
	tr.CreateTransactionRequest.TransactionRequest.TransactionType = transactionType
	return nil
}

func (tr *Transaction) AttachReferenceTransactionId(refTransId string) error {
	tr.CreateTransactionRequest.TransactionRequest.RefTransId = refTransId
	return nil
}

func (tr *Transaction) GetTransactionType() string {
	return tr.CreateTransactionRequest.TransactionRequest.TransactionType
}

func (tr *Transaction) GetTransactionAmount() float64 {
	return tr.CreateTransactionRequest.TransactionRequest.Amount
}
