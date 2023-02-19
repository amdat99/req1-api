

package migrations 


var UsersUp string =  
  `CREATE TABLE "users" (
  "email" VARCHAR(300) PRIMARY KEY NOT NULL,
  "private_id" VARCHAR(50) UNIQUE NOT NULL,
  "id" SERIAL,
  "first_name" VARCHAR(55) DEFAULT '',
  "last_name" VARCHAR(55) DEFAULT '',
  "username" VARCHAR(55) NOT NULL UNIQUE,
  "color" VARCHAR(20),
  "image" VARCHAR(255),
  "email_verified_at" TIMESTAMP,
  "password" VARCHAR(100) NOT NULL,
  "storage_limit" INTEGER DEFAULT 524288000,
  "user_storage_limit" INTEGER DEFAULT 524288000,
  "uploads" JSON,
  "phone" VARCHAR(20),
  "extension_num" VARCHAR(20),
  "disabled" BOOLEAN DEFAULT false,
  "remember_token" VARCHAR(100),
  "otp_secret" VARCHAR(255),
  "last_login" TIMESTAMP,
  "otp_enabled" BOOLEAN DEFAULT false,
  "session_keys" JSONB DEFAULT '[]'::jsonb,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`

var UsersDown string = `DROP TABLE "users";`





