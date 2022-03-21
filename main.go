package main

import (
	"go_gin_crud/routes"
)

func main() {
	//	val = db.SetupDB()
	router := routes.Routes()
	router.Run(":8814")
}
