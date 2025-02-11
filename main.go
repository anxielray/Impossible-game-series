package main

import (
	"os"
	"log"
	m "Zuxyll-game/handlers"
	"github.com/gofiber/fiber/v2"
)

// Global variables
var (
	App   = fiber.New()
	err error
)

func init() {

	err = m.StartDbConn() 
	if err != nil {
		log.Fatal(err)
	}
	m.CreateMoveTable(m.Db)
}


func main() {
	
	//sever the pages
	App.Get("/", func(c *fiber.Ctx) error{
		c.SendFile("public/templates/index.html")
		return nil
	})

	App.Get("/play", func(c *fiber.Ctx) error{
		c.SendFile("public/templates/index.html")
		return nil
	})

	App.Get("/loose", func(c *fiber.Ctx) error{
		c.SendFile("public/templates/loose.html")
		return nil
	})

	App.Get("/win", func(c *fiber.Ctx) error{
		c.SendFile("public/templates/win.html")
		return nil
	})

	// App.Post("/step", m.Game)


	//serve the game
	/* the game terminates when you make a trip and land a step where you were earlier*/

	//define the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1919"
	}

	App.Listen(":"+port)

}
