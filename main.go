package main

import (
	account "github.com/pooraa-nikhil/banking_system_/accounts"
	branch "github.com/pooraa-nikhil/banking_system_/ag"
	customer "github.com/pooraa-nikhil/banking_system_/customer"
	transactions "github.com/pooraa-nikhil/banking_system_/transactions"
	aqua "github.com/rightjoin/aqua"
)

func main() {
	service := aqua.NewRestServer()
	service.AddService(&customer.CustomerService{})
	service.AddService(&transactions.TransactionService{})
	service.AddService(&branch.BranchApi{})
	service.AddService(&account.StartService{})
	service.Run()
}
