
package variableModel

import "req-api/utils"

type VariableType struct {
	Label string `json:"label"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Value string `json:"value" validate:"required"`
	Encrypted bool `json:"encrypted"`
	Additional_data string `json:"additional_data`
	Id int64 `json:"id"`
}

type MultivariableType struct {
	Data []VariableType `json:"data"`
}

//Format variable - Add all records from above struct except the Id and Created_by
func Variable(body *VariableType, session utils.Session ) map[string]interface{} {
	var model =  map[string]interface{}{
		"label": body.Label,
		"org_id": session.Org_id,
		"value": body.Value,
		"encrypted": body.Encrypted,
	}

	return model

}

//Format multi variables
func Multivariable(body *MultivariableType, session utils.Session ) []map[string]interface{} {
	var data []map[string]interface{}
	for _, v := range body.Data {
		data = append(data, Variable(&v, session))
	}
	return data
}

//Fields to query
var Fields = []string{"id","label","created_by","value","encypted","additional_data","view_user_ids","view_team_ids","edit_user_ids","edit_team_ids","delete_user_ids","delete_team_ids","variable_ids","created_at","updated_at",}

