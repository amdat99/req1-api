package migrations

//user email is not referenced in the table as the database where integration is stored may not be the same as the database where the user is stored
var IntegrationUp string = 
`CREATE TABLE integration_1 (
    id SERIAL PRIMARY KEY,
    org_id VARCHAR(50) NOT NULL REFERENCES org_ids(org_id) ON DELETE CASCADE,
	type VARCHAR(50) NOT NULL,
    sub_type VARCHAR(50),
    file BOOLEAN DEFAULT FALSE,
    file_accept VARCHAR(50),
    description VARCHAR(500),
    data JSONB NOT NULL DEFAULT '[]',
    icon VARCHAR(255),
	label VARCHAR(255),
	created_by BIGINT REFERENCES contact_1(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    view_user_ids JSONB DEFAULT '[]'::JSONB,
    view_team_ids JSONB DEFAULT '["all"]'::JSONB,
    edit_user_ids JSONB DEFAULT '[]'::JSONB,
    edit_team_ids JSONB DEFAULT '["all"]'::JSONB,
    delete_user_ids JSONB DEFAULT '[]'::JSONB,
    delete_team_ids JSONB DEFAULT '["all"]'::JSONB
);
`

var IntegrationDown string = `DROP TABLE integration_1;`
