package api

import (
	"github.com/Coflnet/homepage/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func StartWebserver() error {

	engine := html.New("./internal/views", ".html")

	app := fiber.New(fiber.Config{
		AppName: "Coflnet",
		Views:   engine,
	})

	app.Use(logger.New())

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		projects := usecase.ListProjects()

		// render the index template in ../views/
		return c.Render("index", fiber.Map{
			"Projects": projects,
		})
	})

    app.Get("/impressum", func(c *fiber.Ctx) error {
        return c.Render("impressum", fiber.Map{})
    })

    app.Post("/contact", func(c *fiber.Ctx) error {
        firstname := c.FormValue("firstname")
        lastname := c.FormValue("lastname")
        email := c.FormValue("email")
        message := c.FormValue("message")
        err := usecase.SendContactMessage(firstname, lastname, email, message)

        if err != nil {
            return c.Status(200).Render("contact-error", fiber.Map{})
        }

        return c.Render("contact-success", fiber.Map{})
    })

	return app.Listen(":3000")
}
