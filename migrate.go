package main

import (
	"log"
	accounts 		"github.com/pooraa-nikhil/banking_system_/accounts"
	branch 			"github.com/pooraa-nikhil/banking_system_/ag"
	//customer 		"github.com/pooraa-nikhil/banking_system_/customer"
	customer 		"./customer"
	transactions 	"github.com/pooraa-nikhil/banking_system_/transactions"
	lib				"github.com/pooraa-nikhil/banking_system_/lib"
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
	accounts.CreateTables(pg_db)
	branch.Create_Table(pg_db)
}