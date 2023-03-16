package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	submissionController "req-api/controllers/submission"
)

func Submission(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_all":
	 		return submissionController.All(c,route.Session)		
		
		case "POST_single":
			return submissionController.Single(c,route.Session)
		
		case "POST_add":
			return submissionController.Add(c,route.Session)

		case "POST_add-multiple":
			return submissionController.AddMultiple(c,route.Session)

		case "PUT_update":
			return submissionController.Update(c,route.Session)

		case "DELETE_delete":
			return submissionController.Delete(c,route.Session)
		
		case "DELETE_soft-delete":
			return submissionController.SoftDelete(c,route.Session)
		
		case "POST_search":
			return submissionController.Search(c,route.Session)

		case "POST_multi-search":
			return submissionController.MultiSearch(c,route.Session)
			
		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

