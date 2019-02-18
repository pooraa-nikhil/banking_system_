package transactions

import (
	"log"
	"encoding/json"
	"strconv"
	//lib "github.com/pooraa-nikhil/banking_system_/lib"
	lib "github.com/pooraa-nikhil/banking_system_/lib"
	account "github.com/pooraa-nikhil/banking_system_/accounts"
	aqua "github.com/rightjoin/aqua"
)

type TransactionService struct {
	
	aqua.RestService								`root:"transaction" prefix:"service"`
	createTransactionWithdrawal		aqua.POST		`url:"/withdrawal"`
	createTransactionCustToCust		aqua.POST		`url:"/custToCustTransaction"`
	createTransactionDeposit		aqua.POST		`url:"/deposit"`
	getTransactionByType			aqua.GET  		`url:"/getTransactionByType"`

}

func (ts *TransactionService) CreateTransactionWithdrawal(j aqua.Aide) string {

	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp := &Transaction{}
	err := json.Unmarshal([]byte(j.Body) , temp)
	if err != nil {
		log.Printf("Error while Unmarhsalling , Error : %v\n", err)
		return "Failure"
	}
	temp.Type = "debit"

	temp_db, err := pg_db.Begin()

	account := &account.Accounts{}
	account.Account_number = temp.Account_Number
	account.Balance = temp.Amount

	err = account.UpdateSubtractWithTransaction(temp.Account_Number , temp.Amount, temp_db)

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while updating in account, Error : %v\n", err)
		return "Failure"
	}

	err = temp.insertIntoTransaction(temp_db)
	if err != nil {
		temp_db.Rollback()
		log.Printf("error while inserting into transaction, Error : %v\n", err)
		return "failure"
	}

	err = temp_db.Commit()

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while commiting changes, %v\n", err)
		return "Failue"
	}

	return "success"
}

func (ts *TransactionService) CreateTransactionDeposit(j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp := &Transaction{}
	err := json.Unmarshal([]byte(j.Body) , temp)
	if err != nil {
		log.Printf("Error while Unmarhsalling , Error : %v\n", err)
		return "Failure"
	}

	temp.Type = "credit"

	temp_db, err := pg_db.Begin()

	account := &account.Accounts{}
	account.Account_number = temp.Account_Number
	account.Balance = temp.Amount

	err = account.UpdateAddWithTransaction(temp.Account_Number , temp.Amount, temp_db)

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while updating in account, Error : %v\n", err)
		return "Failure"
	}

	err = temp.insertIntoTransaction(temp_db)
	if err != nil {
		log.Printf("error while inserting into transaction, Error : %v\n", err)
		return "Failure"
	}

	err = temp_db.Commit()

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while commiting changes, %v\n", err)
		return "Failue"
	}
	return "success"
}


func (ts *TransactionService) CreateTransactionCustToCust(j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp := &Transaction{}
	err := json.Unmarshal([]byte(j.Body) , temp)
	if err != nil {
		log.Printf("Error while Unmarhsalling , Error : %v\n", err)
		return "Failure"
	}
	if temp.Target_Number == 0 {
		return "Please provide target account number\n"
	}
	temp.Type = "debit"

	temp_db, err := pg_db.Begin()

	account := &account.Accounts{}
	account.Account_number = temp.Account_Number
	account.Balance = temp.Amount
	err = account.UpdateAddWithTransaction(temp.Account_Number , temp.Amount, temp_db)

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while updating in account, Error : %v\n", err)
		return "Failure"
	}

	err = temp.insertIntoTransaction(temp_db)
	if err != nil {
		log.Printf("error while inserting into transaction, Error : %v\n", err)
		temp_db.Rollback()
		return "Failure"
	}

	temp.Account_Number = temp.Target_Number
	temp.Type = "credit"
	temp.Id = 0

	account.Account_number = temp.Account_Number
	err = account.UpdateSubtractWithTransaction(temp.Account_Number , temp.Amount, temp_db)
	if err != nil {
		temp_db.Rollback()
		log.Printf("error while updating in account, Error : %v\n", err)
		return "Failure"
	}

	err = temp.insertIntoTransaction(temp_db)

	if err != nil {
		temp_db.Rollback()
		log.Printf("error while inserting into transaction, Error : %v\n", err)
		return "failure"
	}

	err = temp_db.Commit()

	if err != nil {
		temp_db.Rollback()
		log.Printf("Error while commiting changes, %v\n", err)
		return "Failue"
	}

	return "success"
}


func (ts *TransactionService) GetTransactionByType(j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	acc_no,_ := strconv.Atoi(j.QueryVar["ac_no"])
	var tt = j.QueryVar["type"]
	var trans []Transaction
	trans, err := getTransactionListByType(acc_no,tt,pg_db)
	if err != nil {
		log.Printf("Error while querying in transaction, Error : %v",err)
		return "Failure"
	}

	str, err1 := json.Marshal(trans)
	if err1 != nil {
		log.Printf("Error while marshalling, Error : %v\n", err1)
		return "Failure"
	}

	return string(str)

}