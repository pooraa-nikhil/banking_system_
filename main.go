package main

import (
	//"encoding/json"
	account "github.com/pooraa-nikhil/banking_system_/accounts"
	branch "github.com/pooraa-nikhil/banking_system_/ag"
	customer "github.com/pooraa-nikhil/banking_system_/customer"
	transactions "github/pooraa-nikhil/banking_system_/transactions"
	aqua "github.com/rightjoin/aqua"
)

func main() {
	// customer.CreateTables(pg_db)
	// transactions.CreateTables(pg_db)
	// accounts.CreateTables(pg_db)
	// ag.Create_Table(pg_db)
	service := aqua.NewRestServer()
	service.AddService(&customer.CustomerService{})
	service.AddService(&transactions.TransactionService{})
	service.AddService(&branch.Branch_Api{})
	service.AddService(&accounts.StartService{})
	service.Run()
}
