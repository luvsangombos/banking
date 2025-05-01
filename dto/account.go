package dto

import "github.com/luvsangombos/banking/errs"

type NewAccountResquest struct {
	CustomerId  string  `json:"customerId"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"accountId"`
}

type TransactionRequest struct {
	AccountId       string
	CustomerId      string
	TransactionType string
	Amount          float64
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

func (req TransactionRequest) Validate() *errs.AppError {
	if req.IsTransactionTypeWithdrawal() || req.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Invalid transaction type")
	}

	if req.Amount < 0 {
		return errs.NewValidationError("Invalid amopunt")
	}
	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (req NewAccountResquest) Validate() *errs.AppError {

	if req.Amount < 5000 {
		return errs.NewValidationError("Insuffiecient amount")
	}

	if req.AccountType != "saving" && req.AccountType != "checking" {
		return errs.NewValidationError("Invalid account type")
	}

	return nil

}
