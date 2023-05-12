package tickets

import (
	"api/models"
	"api/utils"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
  tickets := app.Group("/tickets")

  tickets.Get("/types", models.AccountMiddleware, func (c *fiber.Ctx) error {
    ticketTypes, err := models.GetTicketTypes()
    if err != nil {
      return utils.MessageError(c, "Nu s-au putut gasi tipurile de bilete")
    }

    return c.JSON(ticketTypes)
  })


  tickets.Post("/buy-custom", models.AccountMiddleware, func (c *fiber.Ctx) error {
    return c.JSON("ok")
  })
}