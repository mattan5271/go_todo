package main

import (
	"fmt"
	"golang_todo/app/controllers"
	"golang_todo/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
