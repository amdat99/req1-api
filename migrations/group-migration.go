
package migrations 

var GroupUp string =
  `CREATE TABLE "group_1" (
  id SERIAL PRIMARY KEY,
  "org_id" varchar(50) NOT NULL REFERENCES "org_ids" ("org_id") ON DELETE CASCADE,
  "label" varchar(90) NOT NULL,
  "value" varchar(90) NOT NULL,
  "uploads" jsonb DEFAULT '[]'::jsonb,
  "created_by" bigint REFERENCES "contact_1" ("id"),
  "group_lead" bigint REFERENCES "contact_1" ("id"),
  "description" varchar(2000),
  "type" varchar(7) DEFAULT 'group',
  "address" varchar(255),
  "city" varchar(50),
  "state" varchar(50),
  "country" varchar(50),
  "zip" varchar(50),
  "phone" varchar(50),
  "color" varchar(10),


  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp,
  "view_user_ids" jsonb DEFAULT '[]'::jsonb,
  "view_team_ids" jsonb DEFAULT '["all"]'::jsonb,
  "edit_user_ids" jsonb DEFAULT '[]'::jsonb,
  "edit_team_ids" jsonb DEFAULT '["all"]'::jsonb,
  "delete_user_ids" jsonb DEFAULT '[]'::jsonb,
  "delete_team_ids" jsonb DEFAULT '["all"]'::jsonb,
  "workflow_ids" jsonb DEFAULT '[]'::jsonb
);`


var GroupDown string = `DROP TABLE "group_1";`


