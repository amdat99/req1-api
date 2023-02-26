package migrations

var SubmissionUp string = 

`CREATE TABLE submissions_1 (
  id SERIAL PRIMARY KEY,
  org_id VARCHAR(50) NOT NULL REFERENCES org_ids(org_id) ON DELETE CASCADE,
  requirement_id BIGINT NOT NULL REFERENCES requirement_1(id) ON DELETE CASCADE,
  model JSONB DEFAULT '{}'::JSONB,
  label VARCHAR(90),
  value VARCHAR(90),
  uploads JSONB DEFAULT '[]'::JSONB,
  created_by BIGINT REFERENCES contact_1(id),
  has_uploads BOOLEAN DEFAULT false,
  has_tasks BOOLEAN DEFAULT false,
  has_wiki BOOLEAN DEFAULT false,
  "index" FLOAT(10),
  
  additional_fields JSONB DEFAULT '[]'::JSONB,
  additional_values JSONB DEFAULT '[]'::JSONB,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
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

var SubmissionDown string = `DROP TABLE submissions_1;`