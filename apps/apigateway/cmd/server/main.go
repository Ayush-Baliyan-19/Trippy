package main

import (
    "log"
    "net/http"
    "io"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("Apigateway service ok")
    })

    app.Get("/treks", func(c *fiber.Ctx) error {
    resp, err := http.Get("http://trek-service:8001/treks")
    if err != nil { return c.Status(500).SendString("error") }
    body, _ := io.ReadAll(resp.Body)
    return c.Send(body)
})

    log.Fatal(app.Listen(":8000"))
}