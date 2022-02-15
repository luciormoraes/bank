package accounts

import "github.com/luciormoraes/bank/clients"

type Savings struct {
	Owner         clients.Owner
	AgencyNumber  int
	AccountNumber int
	Operation     int
	available     float64
}

func (account *Savings) Withdraw(value float64) string {
	canDo := value <= account.available && value > 0
	if canDo {
		account.available -= value
		return "Ok"
	}
	return "Not enough money"
}

func (account *Savings) Deposit(value float64) (string, float64) {
	if value <= 0 {
		return "Can't Deposit this value", account.available
	}
	account.available += value
	return "Success", account.available
}

func (account *Savings) GetSaldo() float64 {
	return account.available
}
