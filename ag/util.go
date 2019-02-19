package ag

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	account "github.com/pooraa-nikhil/banking_system_/accounts"
	lib "github.com/pooraa-nikhil/banking_system_/lib"
	"time"
)

/*func Conn() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "abcd",
		Database: "bank_pro",
		Addr:     "localhost:5432",
	}
	db := pg.Connect(opts)
	return db
}
*/

// List all entries in the Branch Table
func selectAll() []Branch {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	var br []Branch

	_, err := dbConnect.Query(&br, `select * from branches`)
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}

	return br
}

// List history of all operations of the branch table
func accessHistory() []HistoryBranch {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	var hBr []HistoryBranch

	_, err := dbConnect.Query(&hBr, `select * from history_branches`)
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}

	return hBr
}

// Insert in Branch Table
func insertInTable(data string) error {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	var b = unmarsh(data)

	b.CreatedAt = time.Now()
	b.CreatedBy = "batman"

	inErr := dbConnect.Insert(&b)
	if inErr != nil {
		panic(fmt.Sprintf("%v\n", inErr))
	}

	histErr := insertInHistory(dbConnect, b, "Insert")
	if histErr != nil {
		panic(fmt.Sprintf("%v\n", histErr))
	}

	return nil
}

// Insert in the HistoryBranch Table
func insertInHistory(db *pg.DB, b Branch, ops string) error {

	var hist HistoryBranch

	hist.Info = b
	hist.Operation = ops
	hist.ExecutedBy = "batman"

	if ops == "Insert" {
		hist.Time = b.CreatedAt
	} else {
		hist.Time = time.Now()
	}

	_, histErr := db.Model(&hist).Insert()

	return histErr

}

//Updation in Branch Table
func updateInTable(data string, id int) {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	var upBr = unmarsh(data)

	_, upErr := dbConnect.Model(&upBr).Where("id = ?", id).UpdateNotNull()
	if upErr != nil {
		panic(fmt.Sprintf("%v\n", upErr))
	}

	histErr := insertInHistory(dbConnect, upBr, "Update")
	if histErr != nil {
		panic(fmt.Sprintf("%v\n", histErr))
	}
}

//Delete from Branch Table
func deleteFromTable(data string, id int) string {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	var acc []account.Accounts
	var ac account.Accounts
	var newId int

	idErr := dbConnect.Model((*account.Accounts)(nil)).Column("id").Where("name = ?", "Head Branch").Select(&newId)
	if idErr != nil {
		panic(fmt.Sprintf("%v\n", idErr))
	}

	if newId == id {
		return "cannot delete Head Branch.\n"
	}
	ac.Branch_id = newId

	_, accErr := dbConnect.Query(&acc, `select * from accounts where branch_id = ?`, id)
	if accErr != nil {
		panic(fmt.Sprintf("%v\n", accErr))
	}

	if len(acc) != 0 {
		_, upAccErr := dbConnect.Model(&ac).Where("branch_id = ?", id).UpdateNotNull()
		if upAccErr != nil {
			panic(fmt.Sprintf("%v\n", upAccErr))
		}
	}

	_, delErr := dbConnect.Model(&Branch{}).Where("id = ?", id).Delete()
	if delErr != nil {
		panic(fmt.Sprintf("%v\n", delErr))
	}

	histErr := insertInHistory(dbConnect, Branch{}, "Delete")
	if histErr != nil {
		panic(fmt.Sprintf("%v\n", histErr))
	}

	return "Deletion Successfull \n"
}

//Unmarshal json data
func unmarsh(data string) Branch {

	var byteData = []byte(data)
	var br Branch

	unErr := json.Unmarshal(byteData, &br)
	if unErr != nil {
		//fmt.Println("error in Unmarshal of data in Branch insertion %v\n", un_err)
		panic(fmt.Sprintf("%v\n", unErr))
	}

	return br
}
