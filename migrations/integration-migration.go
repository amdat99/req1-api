package migrations

//user email is not referenced in the table as the database where integration is stored may not be the same as the database where the user is stored
var IntegrationUp string = 
`CREATE TABLE integration_1 (
    id SERIAL PRIMARY KEY,
	integration_key VARCHAR(50) NOT NULL,
    user_email VARCHAR(300) NOT NUll,  
	integration_type VARCHAR(50) NOT NULL,
	label VARCHAR(90),
	created_by BIGINT REFERENCES contact_1(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);
`

var IntegrationDown string = `DROP TABLE integration_1;`