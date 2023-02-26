

package migrations 


var ContactUp string =
  `
CREATE TABLE "contact_1" (
  "id" serial primary key,
  "org_id" varchar(50) not null references "org_ids" ("org_id") on delete cascade,
  "user_id" bigint,
  "label" varchar(90) not null,
  "username" varchar(55) not null,
  "title" varchar(10),
  "first_name" varchar(40) not null,
  "last_name" varchar(40),
  "middle_name" varchar(40),
  "email" varchar(200) not null,
  "alt_email" varchar(200),
  "phone_numb" varchar(20),
  "address" varchar(255),
  "city" varchar(50),
  "zip" varchar(10),
  "state" varchar(50),
  "latitude" float(10),
  "created_by" bigint references "contact_1" ("id"),
  "longitude" float(10),
  "internal" boolean default false,
  "has_wiki" boolean default false,
  "country" varchar(50),
  "color" varchar(10),
  "health_info" jsonb default '[]'::jsonb,
  "contact_type" varchar(20),
  "dob" date,
  "status" varchar(20) default 'Online',  
     
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp,
  "uploads" jsonb default '[]'::jsonb,
  "has_uploads" boolean default false,
  "has_tasks" boolean default false,
  "additional_fields" jsonb default '[]'::jsonb,
  "additional_values" jsonb default '[]'::jsonb,
  "view_user_ids" jsonb default '[]'::jsonb,
  "view_team_ids" jsonb default '["all"]'::jsonb,
  "edit_user_ids" jsonb default '[]'::jsonb,
  "edit_team_ids" jsonb default '["all"]'::jsonb,
  "delete_user_ids" jsonb default '[]'::jsonb,
  "delete_team_ids" jsonb default '["all"]'::jsonb,
  "workflow_ids" jsonb default '[]'::jsonb
);
`

var ContactDown string = `DROP TABLE "contact_1";`



