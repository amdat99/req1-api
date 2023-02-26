package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	reqController "req-api/controllers/requirement"
)

func Requirement(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_all":
	 		return reqController.All(c,route.Session)		
		
		case "POST_single":
			return reqController.Single(c,route.Session)
		
		case "POST_add":
			return reqController.Add(c,route.Session)

		case "POST_add-multiple":
			return reqController.AddMultiple(c,route.Session)

		case "PUT_update":
			return reqController.Update(c,route.Session)

		case "DELETE_delete":
			return reqController.Delete(c,route.Session)
		
		case "DELETE_soft-delete":
			return reqController.SoftDelete(c,route.Session)
		
		case "POST_search":
			return reqController.Search(c,route.Session)

		case "POST_multi-search":
			return reqController.MultiSearch(c,route.Session)
			

		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

