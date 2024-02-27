package main

import (
	"fmt"
	"github.com/JohnKucharsky/real_world_fiber_gorm/db"
	"github.com/JohnKucharsky/real_world_fiber_gorm/handler"
	"github.com/JohnKucharsky/real_world_fiber_gorm/router"
	"github.com/JohnKucharsky/real_world_fiber_gorm/store"
	"log"
)

func main() {
	r := router.New()

	d := db.New()
	err := db.AutoMigrate(d)
	if err != nil {
		log.Fatal("Can't migrate", err.Error())
	}

	us := store.NewUserStore(d)

	h := handler.NewHandler(us)
	h.Register(r)

	err = r.Listen(":8080")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Println("Running cleanup tasks...")
}
