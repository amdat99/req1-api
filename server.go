package main

import (
	"log"
	"req-api/configs"
	"req-api/routes"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()
    
    configs.Init()

    app.Use(limiter.New(limiter.Config{
      Expiration: 1 * time.Second,
      Max: 5,
    }))
    app.Use(logger.New()) // add logger middleware

    app.Get("/", (routes.Message))   
     
    //Http routes 
    app.Use("/auth", routes.Auth)
    app.Use("/org", routes.Org)
    app.Use("/requirement", routes.Requirement)
    
    
    log.Fatal(app.Listen(":5000"))

}


