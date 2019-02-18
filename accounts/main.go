package accounts 

import (

	"log"
	"encoding/json"
	pg "github.com/go-pg/pg"
	aqua "github.com/rightjoin/aqua"


)

var pg_db *pg.DB

type StartService struct {

	aqua.RestService `root:"accounts" prefix:"service"`
	addcust aqua.POST `url:"addCustomer"`
	deletecust aqua.DELETE `url:"deleteCust/{bid:[0-9]+}/{cusid:[0-9]+}"`
	updatecust aqua.PUT `url:"updateCust/{accn:[0-9]+}/{bal:[0-9]+}"`
	updatecustsubtract aqua.PUT `url:"updateCustSubtract/{accn:[0-9]+}/{bal:[0-9]+}"`
//	update aqua.PUT `url:"updateCustomer"`

}

func (u *StartService) Updatecustsubtract(accn int, bal int, j aqua.Aide) string {

	j.LoadVars()
	acc1 := &Accounts{}

	updateErr := acc1.UpdateSubtract(accn,bal,pg_db)

	if updateErr != nil {
		log.Printf("Error while updating. Reason: %v\n",updateErr)
		return "Failure"
	}
	return "Success"

}

func (u *StartService) Updatecust(accn int, bal int, j aqua.Aide) string {

	j.LoadVars()
	acc1 := &Accounts{}

	updateErr := acc1.UpdateAdd(accn,bal,pg_db)

	if updateErr != nil {
		log.Printf("Error while updating. Reason: %v\n",updateErr)
		return "Failure"
	}
	return "Success"

}

func (d *StartService) Deletecust(bid int,cusid int, j aqua.Aide) string {
	
	j.LoadVars()
	acc := &Accounts{}

	deleteErr := acc.DeleteCustomer(bid, cusid, pg_db)

	if deleteErr != nil {
		log.Printf("Error while deleting. Reason: %v\n",deleteErr)
		return "Failure"
	}
	return "Success"
}

func (a *StartService) Addcust(j aqua.Aide) string {
	j.LoadVars()
	emptyObject := &Accounts{}
	err := json.Unmarshal([]byte(j.Body),emptyObject)
	log.Println(emptyObject)
	if err != nil {
		log.Printf("Error occured.\n")
		return "Failure"
	}
	emptyObject.Add(pg_db)
	return "Success"

}

func CreateTables(db *pg.DB) {

	pg_db = db
	CreateAccountsTable(pg_db)
	CreateAccountsHistoryTable(pg_db)

}