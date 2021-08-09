package main

import (
	_ "github.com/go-sql-driver/mysql"
	"oyccx/model"
	"oyccx/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
