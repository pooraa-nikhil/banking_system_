package ag

import (
	//ag "./ag"
	//"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	aqua "github.com/rightjoin/aqua"
)

type Branch_Api struct {
	aqua.RestService `prefix:"bank" root:"branch"`
	innbranch        aqua.POST   `url:"/"`
	updbranch        aqua.PUT    `url:"/{id:[0-9]+}"`
	delbranch        aqua.DELETE `url:"/{id:[0-9]+}"`
	listbranch       aqua.GET    `url:"/"`
	/*listhistory		aqua.GET `url:"/"`*/
}

var db_co *pg.DB

func Create_Table(db *pg.DB) {

	db_co = db

	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	branch_err := db.CreateTable(&Branch{}, opt)

	if branch_err != nil {
		fmt.Printf("error in creation of Branch %v\n", branch_err)
	}

	hist_err := db.CreateTable(&History_Branch{}, opt)

	if hist_err != nil {
		fmt.Printf("error in creation of History_Branch %v\n", hist_err)
	}

}

func (b *Branch_Api) Listbranch(j aqua.Aide) string {
	j.LoadVars()

	branch := Selectall(db_co)

	return fmt.Sprintf("%v\n", branch)

}

func (b *Branch_Api) Innbranch(j aqua.Aide) string {

	j.LoadVars()

	err := Insert_in_Table(db_co, j.Body)

	if err != nil {
		fmt.Println("error")
	}

	return "Insertion in Branch Successfull\n"
}

func (b *Branch_Api) Updbranch(id int, j aqua.Aide) string {
	j.LoadVars()
	Update_in_Table(db_co, j.Body, id)

	return "Updation in Branch Successfull\n"
}

func (b *Branch_Api) Delbranch(id int, j aqua.Aide) string {
	j.LoadVars()
	str := Delete_from_Table(db_co, j.Body, id)

	return str
}

/*func main() {

	db_co = ag.Conn()

	if db_co == nil {
		fmt.Println("Branch : failed to create a connection")
	}

	ag.Create_Table(db_co)

	service := aqua.NewRestServer()
	service.AddService(&Branch_Api{})
	service.Run()
}
*/
