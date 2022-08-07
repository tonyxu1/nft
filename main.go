package main

import (
	"fmt"
	"nft/db"
)

func main() {

	err := db.CreateSchema()
	if err != nil {
		fmt.Println(err)
	}

	startServer()
}
