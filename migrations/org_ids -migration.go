
package migrations 

var OrgIdsUp string =
  `CREATE TABLE "org_ids" (
  "org_id" VARCHAR(50) PRIMARY KEY NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`


var OrgIdsDown string = `DROP TABLE "org_ids";`


