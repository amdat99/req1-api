
package submissionModel

import "req-api/utils"


type SubmissionType struct {
	Label string `json:"label" required:"true"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Id int64 `json:"id"`
	Requirement_id int64 `json:"requirement_id" required:"true"`
	View_id int64 `json:"view_id"`
	Model string `json:"model" required:"true"`
	index int `json:"index"`
}

type MultiSubmissionType struct {
	Data []SubmissionType `json:"data"`
}

//Format Submission - Add all records from above struct except the Id and Created_by
func Submission(body *SubmissionType, session utils.Session ) map[string]interface{} {
	var model =  map[string]interface{}{
		"label": body.Label,
		"org_id": session.Org_id,
		"requirement_id": body.Requirement_id,
		"view_id": body.View_id,
		"model":  utils.JsonArr(body.Model),
	}

	return model

}

//Format multi Submissions
func MultiSubmission(body *MultiSubmissionType, session utils.Session ) []map[string]interface{} {
	var data []map[string]interface{}
	for _, v := range body.Data {
		data = append(data, Submission(&v, session))
	}
	return data
}

//Fields to query
var Fields = []string{"id","requirement_id","view_id","model","label","uploads","created_by","has_uploads","has_tasks","has_wiki","index","additional_fields","additional_values","created_at","updated_at","deleted_at","view_user_ids","view_team_ids","edit_user_ids","edit_team_ids","delete_user_ids","delete_team_ids","workflow_ids"}

