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
		//fmt.Println(err)
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
		return utils.ResError(c, "Error adding requirement", 500)
	}
	return utils.ResAddSuccess(c, requirementData)
}

//Update a requirement
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *requirementModel.RequirementType = new(requirementModel.RequirementType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	updateErr := database.UpdateRow(configs.DB[session.Db], "requirement", session, requirementModel.Requirement(body,session),body.Id)
	if updateErr != nil {
		return utils.ResError(c, "Error updating requirement", 500)
	}
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
