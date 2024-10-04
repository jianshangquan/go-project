package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)


type Message struct {
    Name string
    Body string
    Time int64
}



func main(){
	app := fiber.New();
	app.Get("/", func(c *fiber.Ctx) error {
		// m := Message{"fdsfsd", "fdsfa", 1294706395881547000};

		d := time.Now();
		return c.JSON(d)
	})

	app.Delete("/", func (c *fiber.Ctx) error {
		return c.SendString("Delete");
	});
	app.Listen(":80");
}