
package requirementModel

import "req-api/utils"

type RequirementType struct {
	Label string `json:"label" validate:"required"`
	Type string `json:"type" validate:"required"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Views string `json:"views`
	Additional_data string `json:"additional_data`
	Id int64 `json:"id"`
}

type MultiRequirementType struct {
	Data []RequirementType `json:"data"`
}

//Format requirement - Add all records from above struct except the Id and Created_by
func Requirement(body *RequirementType, session utils.Session ) map[string]interface{} {
	var model =  map[string]interface{}{
		"label": body.Label,
		"type": body.Type,
		"org_id": session.Org_id,
		"views": utils.JsonArr(body.Views),
		"additional_data": utils.JsonArr(body.Additional_data),
	}

	if model["views"] == "[]" {
		model["views"] = "[{\"blocks\":[{\"id\":1,\"subBlocks\":[{\"field\":{},\"id\":1}]}],\"name\":\"Form1\"}]"
	}

	return model

}

//Format multi requirements
func MultiRequirement(body *MultiRequirementType, session utils.Session ) []map[string]interface{} {
	var data []map[string]interface{}
	for _, v := range body.Data {
		data = append(data, Requirement(&v, session))
	}
	return data
}


//Fields to query
var Fields = []string{"id","label","type","created_by","views","additional_data","view_user_ids","view_team_ids","edit_user_ids","edit_team_ids","delete_user_ids","delete_team_ids","workflow_ids","created_at","updated_at"}

