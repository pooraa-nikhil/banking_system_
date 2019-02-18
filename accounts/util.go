package accounts

import (

	"log"
	"errors"
	"time"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//This function returns the details of a customer from Accounts table using the account number.


func (ac *Accounts) GetCustomerByAccn(db *pg.DB, accn int) (Accounts,error) {
	err := db.Model(ac).Where("account_number=?",accn).Select()
	if err != nil {
		return *ac,err
	}
	return *ac,nil
}

//This function adds the balance deposited by the customer to Accounts table using the customers account number.

func (ac *Accounts) UpdateAdd(accn int,bal int,db *pg.DB) error {


	updateErr := db.Model(ac).Column("balance").Where("account_number=?",accn).Select()
	if updateErr != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr)
		return updateErr
	}
	(*ac).Balance = (*ac).Balance + bal
	log.Println((*ac).Balance)

	_,updateErr1 := db.Model(ac).Set("balance=?",(*ac).Balance).Where("account_number=?",accn).Update()
	if updateErr1 != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr1)
		return updateErr1
	}
	log.Printf("The balance was added successfully.")

	custByAccn, err :=ac.GetCustomerByAccn(db,accn)
	if err!= nil {
		log.Printf("There was an error getting the details.Reason: %v\n",err)
		return err
	}
	acc := &Account_history{}
	acc.Accounts = custByAccn
	acc.Operation = "Update"
	acc.Executed_by = "Chaitanya"
	acc.Time = time.Now()
	insertErr := db.Insert(acc)
	if insertErr != nil {
		log.Printf("Error while inserting into Account_history,Reason: %v\n",insertErr)
		return insertErr
	}
	log.Printf("Insertion into Account_history successful")

	return nil

}


//This functions updates the Accounts table with transaction.

func (ac *Accounts) UpdateAddWithTransaction(accn int,bal int,db *pg.Tx) error {

	updateErr := db.Model(ac).Column("balance").Where("account_number=?",accn).Select()
	if updateErr != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr)
		return updateErr
	}
	(*ac).Balance = (*ac).Balance + bal
	log.Println((*ac).Balance)

	_,updateErr1 := db.Model(ac).Set("balance=?",(*ac).Balance).Where("account_number=?",accn).Update()
	if updateErr1 != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr1)
		return updateErr1
	}
	log.Printf("The balance was added successfully.")
	return nil

}


//This function debits balance from Accounts table of a particular customer by using their account number.

func (ac *Accounts) UpdateSubtract(accn int,bal int,db *pg.DB) error {


	log.Println(accn)
	log.Println(bal)
	updateErr := db.Model(ac).Column("balance").Where("account_number=?",accn).Select()
	if updateErr != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr)
		return updateErr
	}
	if (*ac).Balance < bal {
		return errors.New("Insuffient Balance")
	}
	(*ac).Balance = (*ac).Balance - bal
	log.Println((*ac).Balance)

	_,updateErr1 := db.Model(ac).Set("balance=?",(*ac).Balance).Where("account_number=?",accn).Update()
	if updateErr1 != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr1)
		return updateErr1
	}
	custByAccn, err := ac.GetCustomerByAccn(db,accn)
	if err!= nil {
		log.Printf("There was an error getting the details.Reason: %v\n",err)
		return err
	}
	acc := &Account_history{}
	acc.Accounts = custByAccn
	acc.Operation = "Update"
	acc.Executed_by = "Chaitanya"
	acc.Time = time.Now()

	insertErr := db.Insert(acc)
	if insertErr != nil {
		log.Printf("Error while inserting into Account_history,Reason: %v\n",insertErr)
		return insertErr
	}
	log.Printf("Insertion into Account_history successful")

	return nil

}

//This function updates the Accounts table with transaction.

func (ac *Accounts) UpdateSubtractWithTransaction(accn int,bal int,db *pg.Tx) error {

	log.Println(accn)
	log.Println(bal)
	updateErr := db.Model(ac).Column("balance").Where("account_number=?",accn).Select()
	if updateErr != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr)
		return updateErr
	}
	if (*ac).Balance < bal {
		return errors.New("Insuffient Balance")
	}
	(*ac).Balance = (*ac).Balance - bal
	log.Println((*ac).Balance)

	_,updateErr1 := db.Model(ac).Set("balance=?",(*ac).Balance).Where("account_number=?",accn).Update()
	if updateErr1 != nil {
		log.Printf("There was an error updating the balance. Reason: %v\n",updateErr1)
		return updateErr1
	}
	log.Printf("The balance was subtracted successfully.")
	return nil

}

//This function deletes the customer from Accounts table using their branch id and customer id.

func (ac *Accounts) DeleteCustomer(bid int, cusid int,db *pg.DB) error {

	_,deleteErr := db.Model(ac).
					Where("customer_id=?",cusid).
					Where("branch_id=?",bid).Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting the customer. Reason: %v\n",deleteErr)
		return deleteErr
	}
	log.Printf("Customer deleted successfully")
	return nil

}

//This function add a new customer in the Accounts table.

func (ac *Accounts) Add(db *pg.DB) error {

	insertErr := db.Insert(ac)
	if insertErr != nil {
		log.Printf("Error while inserting new customer into DB,Reason: %v\n",insertErr)
		return insertErr
	}
	log.Printf("Insertion successful")
	return nil

}

//This function creates a new Accounts table.

func CreateAccountsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Accounts{}, opts)
	if createErr!=nil {
		log.Printf("Error while creating table Accounts,Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table Accounts created successfully.\n")
	return nil
}

//This function creates a new Account_history table.

func CreateAccountsHistoryTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Account_history{}, opts)
	if createErr!=nil {
		log.Printf("Error while creating table Accounts history,Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table Accounts history created successfully.\n")
	return nil
}

