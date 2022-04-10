/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"CliTaskManager/database"
	"CliTaskManager/pgk"
)

func main() {
	database.ConnectDB()

	pgk.Execute()

	//fmt.Println(string(body))

}
