package accounts

import "github.com/luciormoraes/bank/clients"

type Account struct {
	Owner         clients.Owner
	AgencyNumber  int
	AccountNumber int
	available     float64
}

func (account *Account) Withdraw(value float64) string {
	canDo := value <= account.available && value > 0
	if canDo {
		account.available -= value
		return "Ok"
	}
	return "Not enough money"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value <= 0 {
		return "Can't Deposit this value", account.available
	}
	account.available += value
	return "Success", account.available
}

func (account *Account) Transference(valueToTransfer float64, accountToReceive *Account) bool {
	if valueToTransfer < account.available && valueToTransfer > 0 {
		account.available -= valueToTransfer
		accountToReceive.Deposit(valueToTransfer)
		return true
	}
	return false
}

func (account *Account) GetSaldo() float64 {
	return account.available
}
