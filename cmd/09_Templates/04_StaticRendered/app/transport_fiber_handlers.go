package app

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (app *App) HandlerAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString(
			fmt.Sprintf(
				"articles: %d",
				app.Blog.GetNumberArticles(),
			),
		)
	}
}

func (app *App) HandlerGetArticle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		codeRequest := c.Params("code")

		reconstructedItem, errFetch := app.Blog.GetArticle(codeRequest)
		if errFetch != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{
					"success": false,
					"error":   errFetch.Error(),
				},
			)
		}

		return c.Status(http.StatusOK).JSON(
			&fiber.Map{
				"success": true,
				"article": reconstructedItem,
			},
		)
	}
}
