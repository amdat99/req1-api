
package integrationModel

import "req-api/utils"


type IntegrationType struct {
	Label string `json:"label" validate:"required"`
	Type string `json:"type" validate:"required"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Id int64 `json:"id"`
	Sub_type string `json:"sub_type"`
	File bool `json:"file"`
	File_accept string `json:"file_accept"`
	Description string `json:"description"`
	Data string `json:"data"`
	Icon string `json:"icon"`
}

type MultiIntegrationType struct {
	Data []IntegrationType `json:"data"`
}

//Format Integration - Add all records from above struct except the Id and Created_by
func Integration(body *IntegrationType, session utils.Session ) map[string]interface{} {
	var model =  map[string]interface{}{
		"label": body.Label,
		"type": body.Type,
		"org_id": session.Org_id,
		"sub_type": body.Sub_type,
		"file": body.File,
		"file_accept": body.File_accept,
		"description": body.Description,
		"data": utils.JsonArr(body.Data),
		"icon": body.Icon,

	}

	return model

}

//Format multi Integrations
func MultiIntegration(body *MultiIntegrationType, session utils.Session ) []map[string]interface{} {
	var data []map[string]interface{}
	for _, v := range body.Data {
		data = append(data, Integration(&v, session))
	}
	return data
}


//Fields to query
var Fields = []string{"id","label","type","sub_type","file","file_accept","description","data","icon","created_at","updated_at","created_by","deleted_at"}

