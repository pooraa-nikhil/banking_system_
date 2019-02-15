package customer

import (
	"time"
)


type Customer struct {

	tableName 		struct{}	`sql:"customer"`
	Id				int			`sql:"id,pk,type:serial" json:"id"`
	Name			string		`sql:"name,type:varchar(128)" json:"name"`
	Sex				string		`sql:"sex,type:sex_enum" json:"sex"`
	Email_id		string		`sql:"email_id,type:varchar(128)"json:"email_id`
	Address			string		`sql:"address,type:varchar(512)" json:"address"`
	Contact_number	int			`sql:"contact_number,type:bigint" json:"contact_number"`
	Created_by		string		`sql:"created_by,type:varchar(128)" json:"created_by"`
	Created_At		time.Time   `sql:"created_At" json:"created_at"`

}

type Customer_history struct {

	tableName 		struct{}	`sql:"customer_history"`
	Id   			int			`sql:"id,pk, serial" json:"id"`
	Customer  		Customer    `sql:"customer,type:jsonb" json:"customer"`
	Operation		string		`sql:"operation,type:varchar(128)" json:"operation`
	Executed_by		string		`sql:"executed_by,type:varchar(128)" json:"executed_by"`
	Time 			time.Time 	`sql:"time" json:"time"`

}