package main

import (
	"log"
	"os"
	//"encoding/json"
	ag "./ag"
	pg "github.com/go-pg/pg"
	aqua "github.com/rightjoin/aqua"
	customer "./customer"
	transactions "./transactions"
)

func main() {
	pg_db := Connect()
	if pg_db == nil {
		log.Printf("Error occured while creating database connection\n")
		os.Exit(100)
	}
	log.Printf("Connection Successful\n")


	customer.CreateTables(pg_db)
	transactions.CreateTables(pg_db)
	ag.Create_Table(pg_db)
	service := aqua.NewRestServer()
	service.AddService(&customer.CustomerService{})
	service.AddService(&transactions.TransactionService{})
	service.AddService(&ag.Branch_Api{})
	service.Run()
}

func Connect() (*pg.DB) {
	
	opts := &pg.Options {
		User : "postgres",
		Password : "abcd",
		Addr : "10.1.4.152:5432",
		Database : "bank_pro",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		return nil
	}

	return db
}

