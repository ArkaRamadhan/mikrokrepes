package main

import (
	"fmt"

	"github.com/arkaramadhan/its-vo/common/db"
)

func init() {

	db.LoadEnvVariables()
	db.ConnectToDB()

}

func main() {
	fmt.Println("Hello, World!")
}
