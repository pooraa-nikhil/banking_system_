package main

import (
	//"encoding/json"
	accounts "./accounts"
	ag "./ag"
	customer "./customer"
	transactions "./transactions"
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
	service.AddService(&ag.Branch_Api{})
	service.AddService(&accounts.StartService{})
	service.Run()
}
