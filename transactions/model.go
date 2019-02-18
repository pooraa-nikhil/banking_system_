package transactions

import (
	"time"
)
// 204, 209, 202, 208,201,207,210,200
type Transaction struct {

	tableName 				struct{}		`sql:"transaction"`
	Id						int				`sql:"id,pk,type:serial" json:"id"`
	Target_Number			int				`sql:"-" json:"target_number"`
	Account_Number			int				`sql:"account_number,notnull" json:"account_number"`
	Type 					string			`sql:"ttype,type:type_enum" json:"ttype"`
	Method					string			`sql:"method,type:method_enum" json:"method"`
	Amount					int 			`sql:"amount" json:"amount"`
	Time					time.Time		`sql:"time" json:"time"`

}