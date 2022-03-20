package main

import (
	"backend-food/pkg/infrastucture/db"
	"fmt"
)

func main() {
	_, err := db.NewDB()
	fmt.Print(err)
}
