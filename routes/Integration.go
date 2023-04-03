package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	integrationController"req-api/controllers/integration"
)

func Integration(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_all":
	 		return integrationController.All(c,route.Session)		
		
		case "POST_single":
			return integrationController.Single(c,route.Session)
		
		case "POST_add":
			return integrationController.Add(c,route.Session)

		case "POST_add-multiple":
			return integrationController.AddMultiple(c,route.Session)

		case "PUT_update":
			return integrationController.Update(c,route.Session)

		case "DELETE_delete":
			return integrationController.Delete(c,route.Session)
		
		case "DELETE_soft-delete":
			return integrationController.SoftDelete(c,route.Session)
		
		case "POST_search":
			return integrationController.Search(c,route.Session)

		case "POST_multi-search":
			return integrationController.MultiSearch(c,route.Session)
			

		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

