package router

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/handler"
	"github.com/JohnKucharsky/real_world_fiber_gorm/store"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r *fiber.App, db *pgxpool.Pool) {
	us := store.NewUserStore(db)
	ps := store.NewProductStore(db)
	os := store.NewOrderStore(db)
	h := handler.NewHandler(us, ps, os)

	v1 := r.Group("/api")

	// users
	users := v1.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/", h.GetUsers)
	users.Get("/:id", h.GetOneUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
	// end users

	// products
	products := v1.Group("/products")
	products.Post("/", h.CreateProduct)
	products.Get("/", h.GetProducts)
	products.Get("/:id", h.GetOneProduct)
	products.Put("/:id", h.UpdateProduct)
	products.Delete("/:id", h.DeleteProduct)
	// end products

	// orders
	orders := v1.Group("/orders")
	orders.Post("/", h.CreateOrder)
	orders.Get("/", h.GetOrders)
	orders.Get("/:id", h.GetOneOrder)
	orders.Put("/:id", h.UpdateOrder)
	orders.Delete("/:id", h.DeleteOrder)
	// end orders
}
