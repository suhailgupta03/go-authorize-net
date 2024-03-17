package authorizenet

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	missingRefTransId        = "missing reference transaction id"
	failedRequest            = "request was not successful"
	incorrectTransactionType = "incorrect transaction type"
)

type HTTPRequest struct {
	Test bool
}

func (req *HTTPRequest) transact(transaction *Transaction) (*[]byte, error) {
	url := "https://apitest.authorize.net/xml/v1/request.api"
	if !req.Test {
		url = "https://api.authorize.net/xml/v1/request.api"
	}

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}
	
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(failedRequest + " Returned with status " + resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Authorize net server sends
	// a UTF-8 text string with a Byte Order Mark (BOM).
	// The BOM identifies that the text is UTF-8 encoded,
	// but it should be removed before decoding.
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return &body, nil
}

// ChargeCreditCard Use this method to authorize and capture a credit card payment.
func (req *HTTPRequest) ChargeCreditCard(transaction *Transaction) (*[]byte, error) {
	transactionType := transaction.GetTransactionType()
	if transactionType != TransactionTypeAuthCapture {
		return nil, errors.New(incorrectTransactionType)
	}
	return req.transact(transaction)
}

// AuthCreditCard Use this method to authorize a credit card payment.
// To actually charge the funds you will need to follow up with a capture transaction.
func (req *HTTPRequest) AuthCreditCard(transaction *Transaction) (*[]byte, error) {
	transactionType := transaction.GetTransactionType()
	if transactionType != TransactionTypeAuthOnly {
		return nil, errors.New(incorrectTransactionType)
	}
	return req.transact(transaction)
}

// CaptureAuthedAmount Use this method to capture funds reserved with a
// previous authOnlyTransaction transaction request.
func (req *HTTPRequest) CaptureAuthedAmount(transaction *Transaction, refTransId string) (*[]byte, error) {
	if refTransId == "" {
		return nil, errors.New(missingRefTransId)
	}

	transactionRequest := TransactionRequest{
		TransactionType: TransactionTypePriorAuthCapture,
		Amount:          transaction.GetTransactionAmount(),
		RefTransId:      refTransId,
	}
	transaction.AttachTransactionRequest(transactionRequest)

	return req.transact(transaction)
}

func ToTransactionResponse(body *[]byte) (*Response, error) {
	var response Response
	if err := json.Unmarshal(*body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
