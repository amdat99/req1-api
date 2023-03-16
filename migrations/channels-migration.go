package migrations

var ChannelUp string = 
`CREATE TABLE channel_1 (
    id SERIAL PRIMARY KEY,
    org_id VARCHAR(50) NOT NULL REFERENCES org_ids(org_id) ON DELETE CASCADE,
	channel_id VARCHAR(50) NOT NULL,
	label VARCHAR(90),
	integration_id BIGINT REFERENCES integration_1(id),
	created_by BIGINT REFERENCES contact_1(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);
`

var ChannelDown string = `DROP TABLE channel_1;`