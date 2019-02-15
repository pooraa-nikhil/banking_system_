package transactions

import (
	"os"
	"log"
	"encoding/json"
	"strconv"
	aqua "github.com/rightjoin/aqua"
	pg 	 "github.com/go-pg/pg"	
)

type TransactionService struct {
	aqua.RestService						`root:"transaction" prefix:"service"`
	createTransaction		aqua.POST		`url:"/"`
	getTransactionByType	aqua.GET  		`url:"getTransactionByType"`
}

var pg_db *pg.DB

func CreateTables(db *pg.DB) {
	pg_db = db
	err := CreateTransactionTable(db)
	if err != nil {
		log.Printf("Error while creating table,Error : %v\n", err)
		os.Exit(100)
	}
}

func (transaction *TransactionService) CreateTransaction(j aqua.Aide) string {

	j.LoadVars()
	temp := &Transaction{}
	err := json.Unmarshal([]byte(j.Body) , temp)
	if err != nil {
		log.Printf("Error while Unmarhsalling , Error : %v\n", err)
		return "Failure"
	}


	var str = temp.Method
	switch str {
		case "customer_to_customer" : 
			if temp.Target_Number == 0 {
				return "Please provide target account number\n"
			}

			temp.Type = "debit"

			log.Println(temp)

			err = temp.InsertIntoTransaction(pg_db)
			if err != nil {
				log.Printf("error while inserting into transaction, Error : %v\n", err)
				return "Failure"
			}

			temp.Account_Number = temp.Target_Number
			temp.Type = "credit"
			temp.Id = 0

			log.Println(temp)


			err = temp.InsertIntoTransaction(pg_db)

			if err != nil {
				log.Printf("error while inserting into transaction, Error : %v\n", err)
				return "failure"
			}
		case "cash_withdrawal" :
			temp.Type = "debit"
			err = temp.InsertIntoTransaction(pg_db)
			if err != nil {
				log.Printf("error while inserting into transaction, Error : %v\n", err)
				return "failure"
			}

		case "cash_deposite" :
			temp.Type = "credit"
			err = temp.InsertIntoTransaction(pg_db)
			if err != nil {
				log.Printf("error while inserting into transaction, Error : %v\n", err)
				return "Failure"
			}
		default :
			log.Printf("Please give valid method\n")
			return "failure"
	}
	return "success"
}


func (transaction *TransactionService) GetTransactionByType(j aqua.Aide) string {
	j.LoadVars()

	acc_no,_ := strconv.Atoi(j.QueryVar["ac_no"])
	var tt = j.QueryVar["type"]
	var trans []Transaction
	trans, err := GetTransactionListByType(acc_no,tt,pg_db)
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