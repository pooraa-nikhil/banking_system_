package ag

import (
	"time"
)

type Branch struct {
	Id        int       `sql:"id,pk type:serial"`
	IFSC      string    `sql:"ifsc,unique"`
	Name      string    `sql:"name"`
	Address   string    `sql:"address"`
	CreatedBy string    `sql:"createdby"`
	CreatedAt time.Time `sql:"createdat"`
}

type HistoryBranch struct {
	Info       Branch    `sql:"operation_info"`
	Operation  string    `sql:"operation"`
	ExecutedBy string    `sql:"executedBy"`
	Time       time.Time `sql:"time"`
}
