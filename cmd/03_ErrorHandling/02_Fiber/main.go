package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf(errorMsgValidateInput, errorBadInput)
	}

	if input == badText {
		return fmt.Errorf(
			errorMsgValidateInput,
			&errorBadText{
				badText: input,
			},
		)
	}

	if input == badSpelling {
		return errorBadText{
			badText: input,
		}
	}

	if input == correctText {
		return nil
	}

	return errorApp{
		Area: "Some Area",
		CODE: "101",
		Err:  errors.New("some error"),
	}
}

func main() {
	app := fiber.New()

	app.Get("/:id",
		func(c *fiber.Ctx) error {
			switch c.Params("id") {
			case "1":
				return c.Status(http.StatusBadRequest).JSON(
					&fiber.Map{
						"success": false,
						"error":   validateInput(badInput).Error(),
					},
				)

			case "2":
				return c.Status(http.StatusBadRequest).JSON(
					&fiber.Map{
						"success": false,
						"error":   validateInput(badText).Error(),
					},
				)

			case "3":
				return c.Status(http.StatusBadRequest).JSON(
					&fiber.Map{
						"success": false,
						"error":   validateInput(badSpelling).Error(),
					},
				)

			case "4":
				errValid := validateInput(correctText)
				if errValid != nil {
					return c.Status(http.StatusInternalServerError).JSON(errValid)
				}

				return c.Status(http.StatusOK).JSON(
					&fiber.Map{
						"success": true,
					},
				)

			default:
				return c.JSON(validateInput(""))
			}
		},
	)

	log.Fatal(app.Listen(":3000"))
}
