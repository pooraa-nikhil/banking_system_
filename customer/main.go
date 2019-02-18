package customer

import (
	"time"
	"os"
	"log"
	"encoding/json"
	pg "github.com/go-pg/pg"
	aqua "github.com/rightjoin/aqua"
)

type CustomerService struct {

	pg_db  *pg.DB

	aqua.RestService 					`root:"customer" prefix:"service"`	
	createCustomer 		aqua.POST 		`url:"/"`
	updateCustomer		aqua.PUT		`url:"/{id:[0-9]+}"`
	deleteCustomer		aqua.DELETE 	`url:"/{id:[0-9]+}"`
	getCustomerByID 	aqua.GET 		`url:"/{id:[0-9]+}"`
	listCustomers		aqua.GET 		`url:"/"`

}

var pg_db *pg.DB

func CreateTables(db *pg.DB) {

	pg_db = db
		
	err:= CreateCustomer(pg_db)
	err1 := CreateCustomerHistory(pg_db)
	if err != nil {
		log.Printf("Error while creating table customer, Error : %v\n", err)
		os.Exit(100)
	}

	if err1 != nil {
		log.Printf("Error while creating table customer history, Error : %v\n", err1)
		os.Exit(100)
	}
}

func (service *CustomerService) CreateCustomer(j aqua.Aide) string {

	j.LoadVars()
	temp := &Customer{}
	err := json.Unmarshal([]byte(j.Body), temp)
	if err != nil {
		log.Printf("Error while unmarshalling , Error : %v", err)
		return "Failure"
	}

	temp.Created_At = time.Now()

	log.Printf("%v",temp)

	insertErr := temp.InsertIntoCustomer(pg_db)

	if insertErr != nil {
		log.Printf("Error while inserting into customer, Error : %v\n", insertErr)
		return "failure"
	}

	return "success created"
}

func (service *CustomerService) UpdateCustomer(id int, j aqua.Aide) string {
	j.LoadVars()
	temp := &Customer{}
	err := json.Unmarshal([]byte(j.Body), temp)
	if err != nil {
		log.Printf("Error while unmarshalling, Error : %v\n", err)
		return "Failure"
	}

	updateErr := temp.UpdateIntoCustomer(pg_db, id)
	if updateErr != nil {
		log.Printf("Error while updating customer, Error : %v\n", updateErr)
		return "Failure"
	}

	return "success"
}

func (service *CustomerService) DeleteCustomer(id int , j aqua.Aide) string {
	j.LoadVars()

	deleteErr := DeleteFromCustomer(pg_db, id)
	if deleteErr != nil {
		log.Printf("Error while deleting from customer, Error : %v\n", deleteErr)
		return "failure"
	}

	return "success"
}

func (service *CustomerService) GetCustomerByID(id int,j aqua.Aide) string {
	j.LoadVars()

	temp, err1 := GetCustomerById(pg_db, id)
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

func (service *CustomerService) ListCustomers(j aqua.Aide) string {
	j.LoadVars()

	customers , err := GetAllCustomers(pg_db)
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