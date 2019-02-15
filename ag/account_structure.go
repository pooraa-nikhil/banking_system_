package ag

import (
	"time"
)

type Accounts struct {
	tableName      struct{} `sql:"accounts"`
	Customer_id    int      `sql:"customer_id,not null" json:"customer_id"`
	Branch_id      int      `sql:"branch_id,not null" json:"branch_id"`
	Account_number int      `sql:"account_number,pk" json:"account_number"`
	Account_type   string   `sql:"account_type,type:acc_type" json:"account_type"`
	Balance        int      `sql:"balance,type:bigint" json:"balance"`
}

type Account_history struct {
	Accounts    Accounts  `sql:"accounts,type:jsonb" json:"accounts"`
	Operation   string    `sql:"operation,type:varchar(128)" json:"operation"`
	Executed_by string    `sql:"executed_by,type:varchar(128)" json:"executed_by"`
	Time        time.Time `sql:"time" json:"time"`
}
