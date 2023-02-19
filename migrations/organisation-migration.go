

package migrations 


var OrganizationUp string =
  `CREATE TABLE "organisation" (
  "id" VARCHAR(50) PRIMARY KEY NOT NULL,
  "org_name" VARCHAR(50) UNIQUE NOT NULL,
  "email" VARCHAR(200) UNIQUE NOT NULL,
  "website" VARCHAR(63),
  "address" VARCHAR(150),
  "city" VARCHAR(30),
  "state" VARCHAR(30),
  "country" VARCHAR(30),
  "postal_code" VARCHAR(10),
  "db" INTEGER NOT NULL DEFAULT 1,
  "table_key" INTEGER NOT NULL DEFAULT 1,
  "longitude" FLOAT(10),
  "latitude" FLOAT(10),
  "logo" VARCHAR(150),
  "cover" VARCHAR(150),
  "description" VARCHAR(500),
  "type" VARCHAR(10),
  "status" VARCHAR(10),
  "suspended" BOOLEAN DEFAULT false,
  "storage_limit" BIGINT DEFAULT 10737418240,
  "email_limit" BIGINT DEFAULT 10000,
  "emails_sent" BIGINT DEFAULT 0,
  "map_limit" BIGINT DEFAULT 10000,
  "map_sent" BIGINT DEFAULT 0,
  "storage_used" BIGINT DEFAULT 0,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`


var OrganisationDown string = `DROP TABLE "organisation";`



