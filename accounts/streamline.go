package accounts

type Account struct {
	Owner         string
	AgencyNumber  int
	AccountNumber int
	Available     float64
}

func (account *Account) Withdraw(value float64) string {
	canDo := value <= account.Available && value > 0
	if canDo {
		account.Available -= value
		return "Ok"
	}
	return "Not enough money"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value <= 0 {
		return "Can't Deposit this value", account.Available
	}
	account.Available += value
	return "Success", account.Available
}

func (account *Account) Transference(valueToTransfer float64, accountToReceive *Account) bool {
	if valueToTransfer < account.Available && valueToTransfer > 0 {
		account.Available -= valueToTransfer
		accountToReceive.Deposit(valueToTransfer)
		return true
	}
	return false
}
