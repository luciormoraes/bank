package main

import (
	"fmt"

	"github.com/luciormoraes/bank/accounts"
)

func main() {

	// account1 := Account{Owner: "LM", AgencyNumber: 589, accountNumber: 12, Available: 201.4}
	// account2 := Account{Owner: "LM", AgencyNumber: 589, accountNumber: 12, Available: 201.4}
	// accountB1 := Account{"LM", 589, 12, 201.3}
	// accountB2 := Account{"LM", 589, 12, 201.3}

	// // fmt.Println(account)
	// // fmt.Println(account2)

	// var accountC1 *Account
	// accountC1 = new(Account)
	// accountC1.Owner = "Sal"
	// fmt.Println(&accountC1)

	// var accountC2 *Account
	// accountC2 = new(Account)
	// accountC2.Owner = "Sal"
	// fmt.Println(&accountC2)

	// fmt.Println(account1 == account2)
	// fmt.Println(accountB1 == accountB2)
	// fmt.Println(accountC1 == accountC2)   // addresses are differents
	// fmt.Println(&accountC1 == &accountC2) // looking to address
	// fmt.Println(*accountC1 == *accountC2) // using pointer

	accountLucio := accounts.Account{}
	accountLucio.Owner = "Lucio"
	accountLucio.Available = 500
	fmt.Println(accountLucio.Available)
	fmt.Println(accountLucio.Withdraw(-100))
	fmt.Println(accountLucio.Available)

	// accountLucio.Deposit(1000)
	fmt.Println(accountLucio.Available)
	fmt.Println(accountLucio.Deposit(1000))

	accountSue := accounts.Account{Owner: "Sue", AgencyNumber: 589, AccountNumber: 12, Available: 201.4}
	fmt.Println(accountSue.Owner, accountSue.Available)

	status := accountSue.Transference(300, &accountLucio)
	fmt.Println(status)
	fmt.Println(accountLucio.Owner, accountLucio.Available)
	fmt.Println(accountSue.Owner, accountSue.Available)

}
