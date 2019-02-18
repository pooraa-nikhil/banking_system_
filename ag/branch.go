package ag

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	account ".././accounts"
	"time"
)

type Branch struct {
	Id        int       `sql:"id,pk type:serial"`
	IFSC      string    `sql:"ifsc,unique"`
	Name      string    `sql:"name,notnull"`
	Address   string    `sql:"address,notnull"`
	CreatedBy string    `sql:"createdby"`
	CreatedAt time.Time `sql:"createdat"`
}

type History_Branch struct {
	Info       Branch    `sql:"operation_info"`
	Operation  string    `sql:"operation"`
	ExecutedBy string    `sql:"executedBy"`
	Time       time.Time `sql:"time"`
}

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
func Selectall(db *pg.DB) []Branch {

	var br []Branch

	_, err := db.Query(&br, `select * from branches`)

	if err != nil {
		fmt.Printf("error in selection from branches. %v \n", err)
	}

	return br
}

func Insert_in_Table(db *pg.DB, data string) error {

	var b = Unmarsh(data)

	b.CreatedAt = time.Now()
	b.CreatedBy = "batman"

	//fmt.Println(*b)

	in_err := db.Insert(&b)

	if in_err != nil {
		fmt.Printf("error in insertion in Branch %v \n", in_err)
	}

	herr := Insert_in_history(db, b, "Insert")

	if herr != nil {
		fmt.Printf("error in insertion in History_Branch %v \n", herr)
	}

	return nil

}

func Insert_in_history(db *pg.DB, b Branch, ops string) error {

	var hist History_Branch

	hist.Info = b
	hist.Operation = ops
	hist.ExecutedBy = "batman"

	if ops == "Insert" {
		hist.Time = b.CreatedAt
	} else {
		hist.Time = time.Now()
	}

	_, hist_err := db.Model(&hist).Insert()

	return hist_err

}

func Update_in_Table(db *pg.DB, data string, id int) {

	var up_br = Unmarsh(data)

	_, up_err := db.Model(&up_br).Where("id = ?", id).UpdateNotNull()

	if up_err != nil {
		fmt.Println("error in Updation of Branch %v\n", up_err)
	}

	hist_err := Insert_in_history(db, up_br, "Update")

	if hist_err != nil {
		fmt.Println("error in updation in History_Branch %v\n", hist_err)
	}
}

func Delete_from_Table(db *pg.DB, data string, id int) string {

	/*var del_br = Unmarsh(data)*/

	var acc []account.Accounts
	var ac account.Accounts
	ac.Branch_id = 103

	_, acc_err := db.Query(&acc, `select * from accounts where branch_id = ?`, id)

	if acc_err != nil {
		fmt.Printf("error in retreval from Accounts in Branch Deletion func %v \n", acc_err)
	}

	if len(acc) != 0 {
		//.Set("branch_id = ?", ac.Branch_id)
		_, up_acc_err := db.Model(&ac).Where("branch_id = ?", id).UpdateNotNull()

		if up_acc_err != nil {
			fmt.Printf("error in updation of accounts %v \n", up_acc_err)
		}
	}

	fmt.Println(acc)

	_, del_err := db.Model(&Branch{}).Where("id = ?", id).Delete()

	if del_err != nil {
		fmt.Println("error in deletion from the Branch %v\n", del_err)
	}

	hist_err := Insert_in_history(db, Branch{}, "Delete")

	if hist_err != nil {
		fmt.Println("error in deletion in History_Branch %v\n", hist_err)
	}

	return "Deletion Successfull \n"

}

func Unmarsh(data string) Branch {
	var byte_data = []byte(data)
	var br Branch

	un_err := json.Unmarshal(byte_data, &br)

	if un_err != nil {
		fmt.Println("error in Unmarshal of data in Branch insertion %v\n", un_err)
	}

	return br
}
