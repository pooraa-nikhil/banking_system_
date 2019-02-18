package lib

import(
	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {

	opts := &pg.Options{
		User:     "postgres",
		Password: "abcd",
		Addr:     "10.1.4.152:5432",
		Database: "bank_pro",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		return nil
	}

	return db
}
