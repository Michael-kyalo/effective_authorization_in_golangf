package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {

	app := fiber.New()

	app.Get("/post", handleGetPost)                         //public
	app.Get("/post/manage", onlyAdmin(handleGetPostManage)) //admin

	log.Fatal((app.Listen(":4000")))

}

func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {

		/**
		  this logic can be removed be as complex as possible
		*/
		user := getUserFromDB()
		if user.Role != "admin" {
			return c.Status(403).JSON(fiber.Map{
				"message": "You are not authorized to access this resource",
			})
		}
		return fn(c)
	}
}

func getUserFromDB() User {
	return User{
		Username: "admin",
		Role:     "user",
	}
}

func handleGetPost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

func handleGetPostManage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome Admin",
	})
}
