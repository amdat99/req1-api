package reqController

import (
	"req-api/configs"
	requirementModel "req-api/models/requirement"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

var validate = validator.New()

//Get all requirements
func All(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Paginate(configs.DB[session.Db], "requirement", requirementModel.Fields,c,session)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting requirements", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Get a single requirement
func Single(c *fiber.Ctx, session utils.Session) error {
	
	reqData,err := database.Single(configs.DB[session.Db], "requirement", requirementModel.Fields, c, session)
	if err != nil {
		//fmt.Println(err)
		return utils.ResError(c, "Error getting requirement", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Add a new requirement
func Add(c *fiber.Ctx, session utils.Session) error {
	
	var body *requirementModel.RequirementType = new(requirementModel.RequirementType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	requirementData,err := database.AddRow(configs.DB[session.Db], "requirement", session, requirementModel.Requirement(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding requirement", 500)
	}
	return utils.ResAddSuccess(c, requirementData)
}

//Add multiple requirements
func AddMultiple(c *fiber.Ctx, session utils.Session) error {

	var body *requirementModel.MultiRequirementType = new(requirementModel.MultiRequirementType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	multiReqData,err := database.MultiAddRow(configs.DB[session.Db], "requirement", session, requirementModel.MultiRequirement(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding requirements", 500)
	}
	return utils.ResMultiAddSuccess(c, multiReqData)
}


//Update a requirement
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *requirementModel.RequirementType = new(requirementModel.RequirementType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}
	
	update,err2:= configs.DB[session.Db].Query(fmt.Sprintf("UPDATE requirement_%s SET label=$1, type=$2, views=$3, additional_data=$4 WHERE id=$5 AND org_id=$6", session.Table_key), body.Label, body.Type, body.Views, body.Additional_data, body.Id, session.Org_id)
	if err2 != nil {
		fmt.Println(err2)
		return utils.ResError(c, "Error updating requirement", 500)
	}
	defer update.Close()

	// updateErr := database.UpdateRow(configs.DB[session.Db], "requirement", session, requirementModel.Requirement(body,session),body.Id)
	// if updateErr != nil {
	// 	fmt.Println(updateErr)
	// 	return utils.ResError(c, "Error updating requirement", 500)
	// }
	return utils.ResSuccess(c)
}


//Delete requirements
func Delete(c *fiber.Ctx, session utils.Session) error {

	err := database.DeleteRows(configs.DB[session.Db], "requirement", c, session)
	if err != nil {
		return utils.ResError(c, "Error deleting requirements", 500)
	}
	return utils.ResSuccess(c)
}


//Soft delete requirements
func SoftDelete(c *fiber.Ctx, session utils.Session) error {
	
	err := database.SoftDeleteRows(configs.DB[session.Db], "requirement", c, session)
	if err != nil {
		return utils.ResError(c, "Error soft deleting requirements", 500)
	}
	return utils.ResSuccess(c)
}

//Filter Search 
func Search(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Search(configs.DB[session.Db], "requirement", c, session, database.DefaultFields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting requirements", 500)
	}
	return utils.ResJSON(c, reqData)
}

//Multi Search - Search with multiple filters ( like and equal)
func MultiSearch(c *fiber.Ctx, session utils.Session) error {

	var body *utils.MultiSearchType = new(utils.MultiSearchType)
	utils.ParseBody(c, body)

	reqData,err := database.MultiSearch(configs.DB[session.Db], "requirement", c, session, body, requirementModel.Fields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting requirements", 500)
	}
	return utils.ResJSON(c, reqData)
}
