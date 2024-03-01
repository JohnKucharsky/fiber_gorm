package main

import (
	"fmt"
	"github.com/JohnKucharsky/real_world_fiber_gorm/db"
	"github.com/JohnKucharsky/real_world_fiber_gorm/router"
)

func main() {
	r := router.New()

	d := db.New()

	router.Register(r, d)

	err := r.Listen(":8080")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Println("Running cleanup tasks...")
}
