
package migrations

var GroupUserUp string =
`
CREATE TABLE group_user_1 (
    id SERIAL PRIMARY KEY,
    org_id VARCHAR(50) NOT NULL REFERENCES org_ids(org_id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES contact_1(id) ON DELETE CASCADE,
    group_id BIGINT NOT NULL REFERENCES group_1(id) ON DELETE CASCADE,
    team_name VARCHAR(90) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	deleted_at TIMESTAMP WITH TIME ZONE
);
`

var GroupUserDown string = `DROP TABLE group_user_1;`