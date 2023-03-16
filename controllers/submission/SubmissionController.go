package submissionController

import (
	"req-api/configs"
	submissionModel "req-api/models/submission"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

var validate = validator.New()

//Get all submissions
func All(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Paginate(configs.DB[session.Db], "submission", submissionModel.Fields,c,session)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting submissions", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Get a single submission
func Single(c *fiber.Ctx, session utils.Session) error {
	
	reqData,err := database.Single(configs.DB[session.Db], "submission", submissionModel.Fields, c, session)
	if err != nil {
		//fmt.Println(err)
		return utils.ResError(c, "Error getting submission", 500)
	}
	return utils.ResJSON(c, reqData)
}


//Add a new submission
func Add(c *fiber.Ctx, session utils.Session) error {
	
	var body *submissionModel.SubmissionType = new(submissionModel.SubmissionType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	submissionData,err := database.AddRow(configs.DB[session.Db], "submission", session, submissionModel.Submission(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding submission", 500)
	}
	return utils.ResAddSuccess(c, submissionData)
}

//Add multiple submissions
func AddMultiple(c *fiber.Ctx, session utils.Session) error {

	var body *submissionModel.MultiSubmissionType = new(submissionModel.MultiSubmissionType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	multiReqData,err := database.MultiAddRow(configs.DB[session.Db], "submission", session, submissionModel.MultiSubmission(body,session))
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error adding submissions", 500)
	}
	return utils.ResMultiAddSuccess(c, multiReqData)
}


//Update a submission
func Update(c *fiber.Ctx, session utils.Session) error {

	var body *submissionModel.SubmissionType = new(submissionModel.SubmissionType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}
	
	// _,err2:= configs.DB[session.Db].Query(fmt.Sprintf("UPDATE submission_%s SET label=$1, type=$2, views=$3, additional_data=$4 WHERE id=$5 AND org_id=$6", session.Table_key), body.Label, body.Type, body.Views, body.Additional_data, body.Id, session.Org_id)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return utils.ResError(c, "Error updating submission", 500)
	// }

	updateErr := database.UpdateRow(configs.DB[session.Db], "submission", session, submissionModel.Submission(body,session),body.Id)
	if updateErr != nil {
		fmt.Println(updateErr)
		return utils.ResError(c, "Error updating submission", 500)
	}
	return utils.ResSuccess(c)
}


//Delete submissions
func Delete(c *fiber.Ctx, session utils.Session) error {

	err := database.DeleteRows(configs.DB[session.Db], "submission", c, session)
	if err != nil {
		return utils.ResError(c, "Error deleting submissions", 500)
	}
	return utils.ResSuccess(c)
}


//Soft delete submissions
func SoftDelete(c *fiber.Ctx, session utils.Session) error {
	
	err := database.SoftDeleteRows(configs.DB[session.Db], "submission", c, session)
	if err != nil {
		return utils.ResError(c, "Error soft deleting submissions", 500)
	}
	return utils.ResSuccess(c)
}

//Filter Search 
func Search(c *fiber.Ctx, session utils.Session) error {

	reqData,err := database.Search(configs.DB[session.Db], "submission", c, session, database.DefaultFields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting submissions", 500)
	}
	return utils.ResJSON(c, reqData)
}

//Multi Search - Search with multiple filters ( like and equal)
func MultiSearch(c *fiber.Ctx, session utils.Session) error {

	var body *utils.MultiSearchType = new(utils.MultiSearchType)
	utils.ParseBody(c, body)

	reqData,err := database.MultiSearch(configs.DB[session.Db], "submission", c, session, body, submissionModel.Fields)
	if err != nil {
		fmt.Println(err)
		return utils.ResError(c, "Error getting submissions", 500)
	}
	return utils.ResJSON(c, reqData)
}
