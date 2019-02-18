package accounts

import (

	"log"
	"errors"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*func (achistory *Account_history) list(db *pg.DB) error {

	getErr := db.Model(achistory).Where("account_number=?account_number").Select()
	if getErr != nil {
		log.Printf("Error while getting account history by using account number. Reason: %v\n.", getErr)
		return getErr
	}
	log.Printf("Account history returned successfully.")
	return nil

}*/

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
	return nil

}


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
	log.Printf("The balance was subtracted successfully.")
	return nil

}

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

func (ac *Accounts) Add(db *pg.DB) error {

	insertErr := db.Insert(ac)
	if insertErr != nil {
		log.Printf("Error while inserting new customer into DB,Reason: %v\n",insertErr)
		return insertErr
	}
	log.Printf("Insertion successful")
	return nil

}

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

func CreateAccountsHistoryTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(Account_history{}, opts)
	if createErr!=nil {
		log.Printf("Error while creating table Accounts history,Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table Accounts history created successfully.\n")
	return nil
}