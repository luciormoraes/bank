package main

import (
	"fmt"

	"github.com/luciormoraes/bank/accounts"
	"github.com/luciormoraes/bank/clients"
)

func main() {

	clientSue := clients.Owner{
		Name: "Sue",
		IRD:  "1234512",
		Job:  "Software Engineer",
	}
	accountSue := accounts.Account{Owner: clientSue,
		AgencyNumber:  123,
		AccountNumber: 123456,
		Available:     2000,
	}
	fmt.Println(accountSue)
}
