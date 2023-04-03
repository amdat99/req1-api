package reqController

import (
	"req-api/configs"
	integrationModel "req-api/models/integration"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

var validate = validator.New()

//Get all integrations
func All(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Paginate(configs.DB[session.Db], "integration", integrationModel.Fields,c,session)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting integrations", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Get a single integration
func Single(c *fiber.Ctx, session utils.Session) error {
	
	reqData,err := database.Single(configs.DB[session.Db], "integration", integrationModel.Fields, c, session)
	if err != nil {
		//fmt.Println(err)
		return utils.ResError(c, "Error getting integration", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Add a new integration
func Add(c *fiber.Ctx, session utils.Session) error {
	
	var body *integrationModel.IntegrationType = new(integrationModel.IntegrationType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	integrationData,err := database.AddRow(configs.DB[session.Db], "integration", session, integrationModel.Integration(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding integration", 500)
	}
	return utils.ResAddSuccess(c, integrationData)
}

//Add multiple integrations
func AddMultiple(c *fiber.Ctx, session utils.Session) error {

	var body *integrationModel.MultiIntegrationType = new(integrationModel.MultiIntegrationType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	multiReqData,err := database.MultiAddRow(configs.DB[session.Db], "integration", session, integrationModel.MultiIntegration(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding integrations", 500)
	}
	return utils.ResMultiAddSuccess(c, multiReqData)
}


//Update a integration
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *integrationModel.IntegrationType = new(integrationModel.IntegrationType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}
	
	updateErr := database.UpdateRow(configs.DB[session.Db], "integration", session, integrationModel.Integration(body,session),body.Id)
	if updateErr != nil {
		fmt.Println(updateErr)
		return utils.ResError(c, "Error updating integration", 500)
	}
	return utils.ResSuccess(c)
}


//Delete integrations
func Delete(c *fiber.Ctx, session utils.Session) error {

	err := database.DeleteRows(configs.DB[session.Db], "integration", c, session)
	if err != nil {
		return utils.ResError(c, "Error deleting integrations", 500)
	}
	return utils.ResSuccess(c)
}


//Soft delete integrations
func SoftDelete(c *fiber.Ctx, session utils.Session) error {
	
	err := database.SoftDeleteRows(configs.DB[session.Db], "integration", c, session)
	if err != nil {
		return utils.ResError(c, "Error soft deleting integrations", 500)
	}
	return utils.ResSuccess(c)
}

//Filter Search 
func Search(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Search(configs.DB[session.Db], "integration", c, session, database.DefaultFields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting integrations", 500)
	}
	return utils.ResJSON(c, reqData)
}

//Multi Search - Search with multiple filters ( like and equal)
func MultiSearch(c *fiber.Ctx, session utils.Session) error {

	var body *utils.MultiSearchType = new(utils.MultiSearchType)
	utils.ParseBody(c, body)

	reqData,err := database.MultiSearch(configs.DB[session.Db], "integration", c, session, body, integrationModel.Fields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting integrations", 500)
	}
	return utils.ResJSON(c, reqData)
}
