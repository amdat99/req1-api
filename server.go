package main

import (
	"log"
	"req-api/configs"
	"req-api/routes"
	"time"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"context"
	"fmt"
)

func main() {
    app := fiber.New()
    
    configs.Init()

    app.Use(func(c *fiber.Ctx) error {
   
  c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-With")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Content-Security-Policy", "upgrade-insecure-requests")
	c.Set("Cross-Origin-Resource-Policy", "same-site")
	c.Set("Cross-Origin-Embedder-Policy", "require-corp")
	c.Set("Origin-Agent-Cluster", "?1")
	c.Set("Referrer-Policy", "no-referrer")
	c.Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains; preload")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-XSS-Protection", "0")
	c.Set("x-DNS-Prefetch-Control", "off")
	c.Set("x-download-options", "noopen")
	c.Set("x-permitted-cross-domain-policies", "none")
   if(c.Method() == "OPTIONS") {
        return c.SendStatus(200)
      }

      return c.Next()
    })

	 h := func(c *fiber.Ctx) error {
        sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
        if err := sleepWithContext(c.UserContext(), sleepTime); err != nil {
            return fmt.Errorf("%w: execution error", err)
        }
        return nil
    }


    //app.Use(limiter.New(limiter.Config{
      //Expiration: 1 * time.Second,
     // Max: 5,
    //}))
    app.Use(logger.New()) // add logger middleware

    app.Get("/", (routes.Message))   
     
    //Http routes 
    app.Use("/auth", routes.Auth , timeout.New(h, 2*time.Second))
    app.Use("/org", routes.Org, timeout.New(h, 2*time.Second))
    app.Use("/requirement", routes.Requirement, timeout.New(h, 2*time.Second))
    app.Use("/workflow", routes.Workflow, timeout.New(h, 2*time.Second))
    app.Use("/variable", routes.Variable, timeout.New(h, 2*time.Second))
    app.Use("/submission", routes.Submission, timeout.New(h, 2*time.Second))
    app.Use("/integration", routes.Integration, timeout.New(h, 2*time.Second))
    //app.Use("/generation", routes.Generation, timeout.New(h, 2*time.Second))
    
    
    log.Fatal(app.Listen(":5000"))

}



func sleepWithContext(ctx context.Context, d time.Duration) error {
    timer := time.NewTimer(d)

    select {
    case <-ctx.Done():
        if !timer.Stop() {
            <-timer.C
        }
        return context.DeadlineExceeded
    case <-timer.C:
    }
    return nil
}