package main

import (
	"blogo/route"
	"fmt"
)

func main() {
	fmt.Println("Hello Blogo!")

	router := route.Router()
	err := router.Run()
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}
}
