package lib

import(
	pg "github.com/go-pg/pg"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type conf struct {
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Addr 	 string `yaml:"addr"`
    Database string `yaml:"database"`
}

// function to make a connection to database
func Connect() *pg.DB {

	file := conf{}
	yamlFile, _ := ioutil.ReadFile("conf.yaml")
	_ = yaml.Unmarshal(yamlFile,&file)

	opts := &pg.Options{
		User:     file.Username,
		Password: file.Password,
		Addr:     file.Addr,
		Database: file.Database,
	}

	var db *pg.DB = pg.Connect(opts)

	return db
}
