
package migrations 


var OrgUserUp string = 
  `CREATE TABLE "org-user" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL DEFAULT '',
  "email" VARCHAR(300) NOT NULL REFERENCES "users" ("email") ON DELETE CASCADE,
  "org_id" VARCHAR(50) REFERENCES "organisation" ("id") ON DELETE CASCADE,
  "org_name" VARCHAR(50) DEFAULT '',
  "role" VARCHAR(10) NOT NULL,
  "db" VARCHAR(10) NOT NULL,
  "table_key" VARCHAR(10) NOT NULL,
  "disabled" BOOLEAN DEFAULT false,
  "personal_id" VARCHAR(50) REFERENCES "users" ("private_id") ON DELETE CASCADE,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`

var OrgUserDown string = `DROP TABLE "org-user";`



