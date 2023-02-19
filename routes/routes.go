package routes

import "github.com/gofiber/fiber/v2"

func Message(c *fiber.Ctx) error {
    return c.SendString("👋!")
}
