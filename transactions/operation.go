package transactions

import (
	"log"
	pg  	"github.com/go-pg/pg"
	orm 	"github.com/go-pg/pg/orm"
)

func CreateTransactionTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions {
			IfNotExists : true,
	}
	createError := db.CreateTable(&Transaction{}, opts)
	if createError != nil {
		return createError
	}
	log.Printf("Table successfully created\n")
	return nil
}

func (transation *Transaction) InsertIntoTransaction(db *pg.Tx) error {

	insertError := db.Insert(transation)
	if insertError != nil {
		return insertError
	}
	log.Printf("inserted into transation\n")
	return nil
}

func GetTransactionListByType(acc_no int, tt string, db *pg.DB) ([]Transaction,error) {

	var transactions []Transaction
	log.Println(tt,acc_no)
	err := db.Model(&transactions).Where("account_number=?",acc_no).Where("ttype=?",tt).Select()
	if err != nil {
		return transactions,err
	}
	return transactions,nil
}