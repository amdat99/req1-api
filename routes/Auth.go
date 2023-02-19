package routes

import (
	authcontroller "req-api/controllers/auth"
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

var route = utils.Router(c,false)
   switch route.Name {
	   case "POST_register":
	 		return authcontroller.Register(c)

		case "POST_login":
			return authcontroller.Login(c)
			
	default:
		return utils.ResRouteError(c,route.Name)
	}

}	


func Org(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_enter":
	 		return authcontroller.EnterOrg(c,route.Session)
			
	default:
		return utils.ResRouteError(c,route.Name)
	}

}	


