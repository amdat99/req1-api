
package requirementModel

import "req-api/utils"

type RequirementType struct {
	Label string `json:"label" validate:"required"`
	Type string `json:"type" validate:"required"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Blocks string `json:"blocks`
	Additional_data string `json:"additional_data`
	Id int64 `json:"id"`
}

//Format requirement
func Requirement(body *RequirementType, session utils.Session ) map[string]interface{} {
	return map[string]interface{}{
		"label": body.Label,
		"type": body.Type,
		"org_id": session.Org_id,
		"blocks": utils.JsonArr(body.Blocks),
		"additional_data": utils.JsonArr(body.Additional_data),
	}
}


//Fields to query
var Fields = []string{"id","label","type","created_by","blocks","additional_data","view_user_ids","view_team_ids","edit_user_ids","edit_team_ids","delete_user_ids","delete_team_ids","workflow_ids","created_at","updated_at"}

