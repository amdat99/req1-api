package variableController

import (
	"req-api/configs"
	variableModel "req-api/models/variable"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

var validate = validator.New()

//Get all variables
func All(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Paginate(configs.DB[session.Db], "variable", variableModel.Fields,c,session)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting variables", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Get a single variable
func Single(c *fiber.Ctx, session utils.Session) error {
	
	reqData,err := database.Single(configs.DB[session.Db], "variable", variableModel.Fields, c, session)
	if err != nil {
		//fmt.Println(err)
		return utils.ResError(c, "Error getting variable", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Add a new variable
func Add(c *fiber.Ctx, session utils.Session) error {
	
	var body *variableModel.VariableType = new(variableModel.VariableType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	variableData,err := database.AddRow(configs.DB[session.Db], "variable", session, variableModel.Variable(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding variable", 500)
	}
	return utils.ResAddSuccess(c, variableData)
}

//Add multiple variables
func AddMultiple(c *fiber.Ctx, session utils.Session) error {

	var body *variableModel.MultivariableType = new(variableModel.MultivariableType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	multiReqData,err := database.MultiAddRow(configs.DB[session.Db], "variable", session, variableModel.Multivariable(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding variables", 500)
	}
	return utils.ResMultiAddSuccess(c, multiReqData)
}


//Update a variable
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *variableModel.VariableType = new(variableModel.VariableType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}
	
	// _,err2:= configs.DB[session.Db].Query(fmt.Sprintf("UPDATE variable_%s SET label=$1, type=$2, views=$3, additional_data=$4 WHERE id=$5 AND org_id=$6", session.Table_key), body.Label, body.Type, body.Views, body.Additional_data, body.Id, session.Org_id)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return utils.ResError(c, "Error updating variable", 500)
	// }

	updateErr := database.UpdateRow(configs.DB[session.Db], "variable", session, variableModel.Variable(body,session),body.Id)
	if updateErr != nil {
		fmt.Println(updateErr)
		return utils.ResError(c, "Error updating variable", 500)
	}
	return utils.ResSuccess(c)
}


//Delete variables
func Delete(c *fiber.Ctx, session utils.Session) error {

	err := database.DeleteRows(configs.DB[session.Db], "variable", c, session)
	if err != nil {
		return utils.ResError(c, "Error deleting variables", 500)
	}
	return utils.ResSuccess(c)
}


//Soft delete variables
func SoftDelete(c *fiber.Ctx, session utils.Session) error {
	
	err := database.SoftDeleteRows(configs.DB[session.Db], "variable", c, session)
	if err != nil {
		return utils.ResError(c, "Error soft deleting variables", 500)
	}
	return utils.ResSuccess(c)
}

//Filter Search 
func Search(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Search(configs.DB[session.Db], "variable", c, session, database.DefaultFields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting variables", 500)
	}
	return utils.ResJSON(c, reqData)
}

//Multi Search - Search with multiple filters ( like and equal)
func MultiSearch(c *fiber.Ctx, session utils.Session) error {

	var body *utils.MultiSearchType = new(utils.MultiSearchType)
	utils.ParseBody(c, body)

	reqData,err := database.MultiSearch(configs.DB[session.Db], "variable", c, session, body, variableModel.Fields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting variables", 500)
	}
	return utils.ResJSON(c, reqData)
}
