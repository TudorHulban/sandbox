package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func sendResponse(c *fiber.Ctx, rawResponse string) error {
	response := jsonAnswer{
		Response: rawResponse,
	}

	return json.NewEncoder(c).Encode(&response)
}
