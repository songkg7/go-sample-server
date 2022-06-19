package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	framework string
	port      = flag.Int("p", 8888, "서버가 Listen할 port 번호를 입력해주세요.")
)

func init() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("하나의 인자를 전달해 framework 를 정의해주세요.")
	}
	framework = flag.Arg(0)
}

func main() {
	switch framework {
	case "fiber":
		RunNewFiberServer()
	}
}

func RunNewFiberServer() {
	addr := fmt.Sprintf(":%d", *port)
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})
	log.Printf("Server is listening %d", *port)
	if err := app.Listen(addr); err != nil {
		log.Print(err)
	}
}
