package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	generationController "req-api/controllers/generation"
)

func Generation(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_gen":
	 		return generationController.Gen(c,route.Session)		
		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

