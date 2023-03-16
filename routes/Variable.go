package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	variableController "req-api/controllers/variable"
)

func Variable(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_all":
	 		return variableController.All(c,route.Session)		
		
		case "POST_single":
			return variableController.Single(c,route.Session)
		
		case "POST_add":
			return variableController.Add(c,route.Session)

		case "POST_add-multiple":
			return variableController.AddMultiple(c,route.Session)

		case "PUT_update":
			return variableController.Update(c,route.Session)

		case "DELETE_delete":
			return variableController.Delete(c,route.Session)
		
		case "DELETE_soft-delete":
			return variableController.SoftDelete(c,route.Session)
		
		case "POST_search":
			return variableController.Search(c,route.Session)

		case "POST_multi-search":
			return variableController.MultiSearch(c,route.Session)
			

		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

