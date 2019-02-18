package customer

import (
	ag ".././accounts"
	"errors"
	pg "github.com/go-pg/pg"

	orm "github.com/go-pg/pg/orm"
	"log"
	"time"
)

func CreateCustomer(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&Customer{}, opts)
	if createErr != nil {
		return createErr
	}
	log.Printf("customer table created\n")
	return nil
}

func (customer *Customer) InsertIntoCustomer(db *pg.DB) error {

	insertError := db.Insert(customer)
	if insertError != nil {
		return insertError
	}

	log.Printf("Added to customer\n")
	return nil
}

func (customer *Customer) UpdateIntoCustomer(db *pg.DB, id int) error {



	_ , updateErr := db.Model(customer).Where("id=?", id).Returning("*").UpdateNotNull()
	
	log.Printf("%v\n", customer)

	if updateErr != nil {
		return updateErr
	}

	custById, err := GetCustomerById(db,id)


	if err != nil {
		return err
	}

	cust := Customer_history{}
	cust.Customer = custById
	cust.Operation = "update"
	cust.Executed_by = "Nikhil"
	cust.Time = time.Now()

	log.Printf("%v\n", cust)

	cust.InsertIntoCustomerHistroy(db)


	log.Printf("Success update\n")
	return nil
}

func DeleteFromCustomer(db *pg.DB, id int) error {

	rows, err := db.Model((*ag.Accounts)(nil)).Where("customer_id=?", id).Count()
	log.Println(rows)

	if rows > 0 {
		return errors.New("Delete accounts of this customer first\n")
	}

	// TODO: forign key constraint from account table


	var customer Customer
	_ , deleteErr := db.Model(&customer).Where("id=?",id).Returning("*").Delete()

	log.Printf("%v\n", customer)
	if deleteErr != nil {
		return deleteErr
	}

	custById, err := GetCustomerById(db,id)


	if err != nil {
		return err
	}

	cust := Customer_history{}
	cust.Customer = custById
	cust.Operation = "delete"
	cust.Executed_by = "Nikhil"
	cust.Time = time.Now()

	log.Printf("%v\n", cust)

	cust.InsertIntoCustomerHistroy(db)


	log.Printf("Success delete\n")
	return nil
}

func GetCustomerById(db *pg.DB, id int) (Customer, error) {
	customer := &Customer{}
	err := db.Model(customer).Where("id=?", id).Select()
	if err != nil {
		return *customer, err
	}
	return *customer, nil
}

func GetAllCustomers(db *pg.DB) ([]Customer, error) {
	var customer []Customer
	err := db.Model(&customer).Select()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// ======== Customer History ===========

func CreateCustomerHistory(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&Customer_history{}, opts)
	if createErr != nil {
		return createErr
	}
	log.Printf("customer history table created\n")
	return nil
}

func (history *Customer_history) InsertIntoCustomerHistroy(db *pg.DB) {

	insertError := db.Insert(history)
	if insertError != nil {
		log.Printf("Error : %v", insertError)
		return
	}
	log.Printf("added to customer history")
	return
}

func DeleteFromCustomerHistory(db *pg.DB, id int) error {

	// TODO: forign key constraint from account table

	var customer Customer_history
	_, deleteErr := db.Model(&customer).Where("id=?", id).Delete()

	log.Printf("%v\n", customer)
	if deleteErr != nil {
		return deleteErr
	}

	log.Printf("Success delete\n")
	return nil
}
