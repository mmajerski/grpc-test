package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/userq11/grpc-test/api"
	"xorm.io/xorm"
)

func main() {
	db, err := xorm.NewEngine("mysql", "root:secret@tcp(localhost:3306)/grpc")
	if err != nil {
		panic(err)
	}

	_, err = db.DBMetas()
	if err != nil {
		panic(err)
	}

	api.Run(9090, db)
}
