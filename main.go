package main

import (
	"api/accounts"
	"api/db"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var RedisOptions *redis.Options = &redis.Options {
  Addr: "127.0.0.1:6379",
  Password: "",
  DB: 0,
}

func main() {
  app := fiber.New()

  db.InitDB()
  db.InitCache(RedisOptions)

  app.Get("/test", func (c *fiber.Ctx) error {
    return c.SendString("it's working")
  })

  accounts.Routes(app)

  app.Listen(":4200")
}