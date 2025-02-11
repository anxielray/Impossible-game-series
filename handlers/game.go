package handlers

import (
	"fmt"
	// "strconv"
	"github.com/gofiber/fiber/v2"
)



var AnsArray = Answers{
	Ans: []Answer{{
		First:  0,
		Second: 0,
		Third:  0,
		Fourth: 1,
	}},
}

func Game(c *fiber.Ctx) error {

	
	var Ans Answer

	if Ans.First == 0 && Ans.Second == 0 && Ans.Third == 0 && Ans.Fourth == 0 {

		App.Get("/win", func(c *fiber.Ctx) error{
			c.SendFile("public/templates/win.html")
			return nil	
		})

	} else {

		fmt.Println("redoing the level")
		App.Get("/", func(c *fiber.Ctx) error{
			c.SendFile("public/templates/index.html")
            return nil
		})
	}
	return nil
}
