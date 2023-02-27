package migrations

var WorkflowUp string = 
`CREATE TABLE workflow_1 (
    id SERIAL PRIMARY KEY,
    org_id VARCHAR(50) NOT NULL REFERENCES org_ids(org_id) ON DELETE CASCADE,
    label VARCHAR(90) NOT NULL,
    type VARCHAR(18) NOT NULL,
    nodes JSONB DEFAULT '[]'::JSONB,
    edges JSONB DEFAULT '[]'::JSONB,
    additional_data JSONB DEFAULT '[]'::JSONB,
    
	created_by BIGINT REFERENCES contact_1(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    view_user_ids JSONB DEFAULT '[]'::JSONB,
    view_team_ids JSONB DEFAULT '["all"]'::JSONB,
    edit_user_ids JSONB DEFAULT '[]'::JSONB,
    edit_team_ids JSONB DEFAULT '["all"]'::JSONB,
    delete_user_ids JSONB DEFAULT '[]'::JSONB,
    delete_team_ids JSONB DEFAULT '["all"]'::JSONB,
    workflow_ids JSONB DEFAULT '[]'::JSONB
);
`

var WorkflowDown string = `DROP TABLE workflow_1;`