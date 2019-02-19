package main

import (
	accounts "github.com/pooraa-nikhil/banking_system_/accounts"
	branch "github.com/pooraa-nikhil/banking_system_/ag"
	"log"
	//customer 		"github.com/pooraa-nikhil/banking_system_/customer"
	customer "./customer"
	lib "github.com/pooraa-nikhil/banking_system_/lib"
	transactions "github.com/pooraa-nikhil/banking_system_/transactions"
)

func main() {

	pg_db := lib.Connect()

	if pg_db == nil {
		log.Printf("Unable to connect database\n")
		panic(100)
	}

	err := customer.CreateCustomer(pg_db)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	err = customer.CreateCustomerHistory(pg_db)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	err = transactions.CreateTransactionTable(pg_db)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}

	err := accounts.CreateAccountsTable(pg_db)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}

	err := accounts.CreateAccountsHistoryTable(pg_db)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}

	branch.Create_Table(pg_db)
}
