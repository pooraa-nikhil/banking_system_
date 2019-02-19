package ag

import (
	"fmt"
	lib "github.com/pooraa-nikhil/banking_system_/lib"
	//"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	aqua "github.com/rightjoin/aqua"
)

type BranchApi struct {
	aqua.RestService `prefix:"bank" root:"branch"`
	innBranch        aqua.POST   `url:"/"`
	updBranch        aqua.PUT    `url:"/{id:[0-9]+}"`
	delBranch        aqua.DELETE `url:"/{id:[0-9]+}"`
	listBranch       aqua.GET    `url:"/"`
	/*listhistory		aqua.GET `url:"/"`*/
}

// Creation of Branch and HistoryBranch Tables
func CreateTable() {

	dbConnect := lib.Connect()
	defer dbConnect.Close()

	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	branchErr := dbConnect.CreateTable(&Branch{}, opt)

	if branchErr != nil {
		fmt.Printf("error in creation of Branch %v\n", branchErr)
		panic(fmt.Sprintf("%v\n", branchErr))
	}

	histErr := dbConnect.CreateTable(&HistoryBranch{}, opt)

	if histErr != nil {
		panic(fmt.Sprintf("%v\n", histErr))
	}

}

//API for Listing all Branch of the Bank
func (b *BranchApi) ListBranch(j aqua.Aide) string {
	j.LoadVars()

	branch := selectAll()

	return fmt.Sprintf("%v\n", branch)

}

// API for insertion of new Branch
func (b *BranchApi) InnBranch(j aqua.Aide) string {

	j.LoadVars()

	err := insertInTable(j.Body)

	if err != nil {
		fmt.Println("error")
		panic(fmt.Sprintf("%v\n", err))
	}

	return "Insertion in Branch Successfull\n"
}

//API for Updation of existing Branch
func (b *BranchApi) UpdBranch(id int, j aqua.Aide) string {
	j.LoadVars()
	updateInTable(j.Body, id)

	return "Updation in Branch Successfull\n"
}

//API for deletion for of an existing Branch
func (b *BranchApi) DelBranch(id int, j aqua.Aide) string {
	j.LoadVars()
	str := deleteFromTable(j.Body, id)

	return str
}

/*
func main() {

		db_co = ag.Conn()

		if db_co == nil {
			fmt.Println("Branch : failed to create a connection")
		}

	ag.Create_Table(db_co)

	service := aqua.NewRestServer()
	service.AddService(&BranchApi{})
	service.Run()
}
*/
