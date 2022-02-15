package main

import (
	"fmt"

	"github.com/luciormoraes/bank/accounts"
	"github.com/luciormoraes/bank/clients"
)

func PayBill(account verifyAccountType, billValue float64) {
	account.Withdraw(billValue)
}

type verifyAccountType interface {
	Withdraw(value float64) string
}

func main() {

	clientSue := clients.Owner{
		Name: "Sue",
		IRD:  "1234512",
		Job:  "Software Engineer",
	}
	accountSue := accounts.Account{Owner: clientSue,
		AgencyNumber:  123,
		AccountNumber: 123456,
	}
	accountSue.Deposit(1000)
	fmt.Println(accountSue)
	fmt.Println(accountSue.GetSaldo())

	savingSue := accounts.Savings{Owner: clientSue, AgencyNumber: 123, AccountNumber: 654321}
	savingSue.Deposit(3000)
	fmt.Println(savingSue)
	savingSue.Withdraw(4000)

	PayBill(&accountSue, 600)
	fmt.Println("Available after payment:", accountSue.GetSaldo())

	PayBill(&savingSue, 600)
	fmt.Println("Available after payment:", savingSue.GetSaldo())

}
