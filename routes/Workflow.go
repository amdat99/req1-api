package routes

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	workflowController "req-api/controllers/workflow"
)

func Workflow(c *fiber.Ctx) error {

var route = utils.Router(c,true)
   switch route.Name{
	   case "POST_all":
	 		return workflowController.All(c,route.Session)		
		
		case "POST_single":
			return workflowController.Single(c,route.Session)
		
		case "POST_add":
			return workflowController.Add(c,route.Session)

		case "POST_add-multiple":
			return workflowController.AddMultiple(c,route.Session)

		case "PUT_update":
			return workflowController.Update(c,route.Session)

		case "DELETE_delete":
			return workflowController.Delete(c,route.Session)
		
		case "DELETE_soft-delete":
			return workflowController.SoftDelete(c,route.Session)
		
		case "POST_search":
			return workflowController.Search(c,route.Session)

		case "POST_multi-search":
			return workflowController.MultiSearch(c,route.Session)
			

		default:
			return utils.ResRouteError(c,route.Name)
		}

}	

