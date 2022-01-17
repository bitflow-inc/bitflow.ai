package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

/*

	"github.com/markbates/pkger"
	"log"
*/

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"tab1": "on",
		})
	})

	app.Get("/solution", func(c *fiber.Ctx) error {
		return c.Render("solution", fiber.Map{
			"tab2": "on",
		})
	})

	app.Get("/press", func(c *fiber.Ctx) error {
		return c.Render("press", fiber.Map{
			"tab3": "on",
		})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"tab4": "on",
		})
	})

	app.Static("/", "./static")

	app.Listen(":3000")
}
