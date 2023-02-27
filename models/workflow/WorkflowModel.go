
package workflowModel

import "req-api/utils"

type WorkflowType struct {
	Label string `json:"label" validate:"required"`
	Type string `json:"type" validate:"required"`
	Org_id string `json:"org_id"`
	Created_by int `json:"created_by"`
	Nodes string `json:"nodes`
	Edges string `json:"edges`
	Additional_data string `json:"additional_data`
	Id int64 `json:"id"`
}

type MultiWorkflowType struct {
	Data []WorkflowType `json:"data"`
}

//Format workflow - Add all records from above struct except the Id and Created_by
func Workflow(body *WorkflowType, session utils.Session ) map[string]interface{} {
	var model =  map[string]interface{}{
		"label": body.Label,
		"type": body.Type,
		"org_id": session.Org_id,
		"edges": utils.JsonArr(body.Edges),
		"nodes": utils.JsonArr(body.Nodes),
		"additional_data": utils.JsonArr(body.Additional_data),
	}

	return model

}

//Format multi workflows
func MultiWorkflow(body *MultiWorkflowType, session utils.Session ) []map[string]interface{} {
	var data []map[string]interface{}
	for _, v := range body.Data {
		data = append(data, Workflow(&v, session))
	}
	return data
}

//Fields to query
var Fields = []string{"id","label","type","created_by","nodes","edges","additional_data","view_user_ids","view_team_ids","edit_user_ids","edit_team_ids","delete_user_ids","delete_team_ids","workflow_ids","created_at","updated_at"}

