package workflowController

import (
	"req-api/configs"
	workflowModel "req-api/models/workflow"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

var validate = validator.New()

//Get all workflows
func All(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Paginate(configs.DB[session.Db], "workflow", workflowModel.Fields,c,session)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting workflows", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Get a single workflow
func Single(c *fiber.Ctx, session utils.Session) error {
	
	reqData,err := database.Single(configs.DB[session.Db], "workflow", workflowModel.Fields, c, session)
	if err != nil {
		//fmt.Println(err)
		return utils.ResError(c, "Error getting workflow", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Add a new workflow
func Add(c *fiber.Ctx, session utils.Session) error {
	
	var body *workflowModel.WorkflowType = new(workflowModel.WorkflowType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	workflowData,err := database.AddRow(configs.DB[session.Db], "workflow", session, workflowModel.Workflow(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding workflow", 500)
	}
	return utils.ResAddSuccess(c, workflowData)
}

//Add multiple workflows
func AddMultiple(c *fiber.Ctx, session utils.Session) error {

	var body *workflowModel.MultiWorkflowType = new(workflowModel.MultiWorkflowType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	multiReqData,err := database.MultiAddRow(configs.DB[session.Db], "workflow", session, workflowModel.MultiWorkflow(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding workflows", 500)
	}
	return utils.ResMultiAddSuccess(c, multiReqData)
}


//Update a workflow
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *workflowModel.WorkflowType = new(workflowModel.WorkflowType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}
	
	// _,err2:= configs.DB[session.Db].Query(fmt.Sprintf("UPDATE workflow_%s SET label=$1, type=$2, views=$3, additional_data=$4 WHERE id=$5 AND org_id=$6", session.Table_key), body.Label, body.Type, body.Views, body.Additional_data, body.Id, session.Org_id)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return utils.ResError(c, "Error updating workflow", 500)
	// }

	updateErr := database.UpdateRow(configs.DB[session.Db], "workflow", session, workflowModel.Workflow(body,session),body.Id)
	if updateErr != nil {
		fmt.Println(updateErr)
		return utils.ResError(c, "Error updating workflow", 500)
	}
	return utils.ResSuccess(c)
}


//Delete workflows
func Delete(c *fiber.Ctx, session utils.Session) error {

	err := database.DeleteRows(configs.DB[session.Db], "workflow", c, session)
	if err != nil {
		return utils.ResError(c, "Error deleting workflows", 500)
	}
	return utils.ResSuccess(c)
}


//Soft delete workflows
func SoftDelete(c *fiber.Ctx, session utils.Session) error {
	
	err := database.SoftDeleteRows(configs.DB[session.Db], "workflow", c, session)
	if err != nil {
		return utils.ResError(c, "Error soft deleting workflows", 500)
	}
	return utils.ResSuccess(c)
}

//Filter Search 
func Search(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Search(configs.DB[session.Db], "workflow", c, session, database.DefaultFields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting workflows", 500)
	}
	return utils.ResJSON(c, reqData)
}

//Multi Search - Search with multiple filters ( like and equal)
func MultiSearch(c *fiber.Ctx, session utils.Session) error {

	var body *utils.MultiSearchType = new(utils.MultiSearchType)
	utils.ParseBody(c, body)

	reqData,err := database.MultiSearch(configs.DB[session.Db], "workflow", c, session, body, workflowModel.Fields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting workflows", 500)
	}
	return utils.ResJSON(c, reqData)
}
