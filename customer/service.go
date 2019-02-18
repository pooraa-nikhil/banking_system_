package customer

import (
	"time"
	"log"
	"encoding/json"
	//pg "github.com/go-pg/pg"
	lib "../lib"
	aqua "github.com/rightjoin/aqua"
)

type CustomerService struct {

	aqua.RestService 					`root:"customer" prefix:"service"`	
	createCustomer 		aqua.POST 		`url:"/"`
	updateCustomer		aqua.PUT		`url:"/{id:[0-9]+}"`
	deleteCustomer		aqua.DELETE 	`url:"/{id:[0-9]+}"`
	getCustomerByID 	aqua.GET 		`url:"/{id:[0-9]+}"`
	listCustomers		aqua.GET 		`url:"/"`

}

// func Init() {

// 	pg_db = pg.Connect()
		
// 	err:= CreateCustomer(pg_db)
// 	err1 := CreateCustomerHistory(pg_db)
// 	if err != nil {
// 		log.Printf("Error while creating table customer, Error : %v\n", err)
// 		os.Exit(100)
// 	}

// 	if err1 != nil {
// 		log.Printf("Error while creating table customer history, Error : %v\n", err1)
// 		os.Exit(100)
// 	}
// }

func (cs *CustomerService) CreateCustomer(j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp := &Customer{}
	err := json.Unmarshal([]byte(j.Body), temp)
	if err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return "Failure"
	}

	temp.Created_At = time.Now()

	log.Printf("%v",temp)

	insertErr := temp.insertIntoCustomer(pg_db)

	if insertErr != nil {
		log.Printf("Error while inserting into customer, Error : %v\n", insertErr)
		return "failure"
	}

	return "success"
}

func (cs *CustomerService) UpdateCustomer(id int, j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp := &Customer{}
	err := json.Unmarshal([]byte(j.Body), temp)
	if err != nil {
		log.Printf("Error while unmarshalling, Error : %v\n", err)
		return "Failure"
	}

	updateErr := temp.updateIntoCustomer(pg_db, id)
	if updateErr != nil {
		log.Printf("Error while updating customer, Error : %v\n", updateErr)
		return "Failure"
	}

	return "success"
}

func (cs *CustomerService) DeleteCustomer(id int , j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	deleteErr := deleteFromCustomer(pg_db, id)
	if deleteErr != nil {
		log.Printf("Error while deleting from customer, Error : %v\n", deleteErr)
		return "failure"
	}

	return "success"
}

func (cs *CustomerService) GetCustomerByID(id int,j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	temp, err1 := getCustomerById(pg_db, id)
	if err1 != nil {
		log.Printf("Error while querying by id, Error : \n",err1)
		return "failure"
	}
	resp , marshError := json.Marshal(temp)
	if marshError != nil {
		log.Printf("Error while marshalling data, Error : %v\n", marshError)
		return "failure"
	} 
	return string(resp)
}

func (cs *CustomerService) ListCustomers(j aqua.Aide) string {
	
	pg_db := lib.Connect()
	defer pg_db.Close()

	j.LoadVars()
	customers , err := getAllCustomers(pg_db)
	if err != nil {
		log.Printf("Error while querying by id\n")
		return "failure"
	}
	resp , marshError := json.Marshal(customers)
	if marshError != nil {
		log.Printf("Error while marshalling data, Error : %v\n", marshError)
		return "failure"
	} 
	return string(resp)
	
}