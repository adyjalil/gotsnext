CREATE TABLE "users" (
  "username" varchar(50) PRIMARY KEY,
  "email" varchar(50) UNIQUE NOT NULL,
  "password" varchar(100) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "role" (
  "name" varchar(50) PRIMARY KEY,
  "description" varchar(100) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "team" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "todo" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(100) NOT NULL,
  "start_time" timestamp,
  "end_time" timestamp,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "team_members" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "team_id" varchar(50) NOT NULL,
  "user_id" varchar(50) NOT NULL,
  "role" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "todo_actor" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "todo_id" varchar(50) NOT NULL,
  "user_id" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp
);

ALTER TABLE "team_members" ADD FOREIGN KEY ("team_id") REFERENCES "team" ("id");

ALTER TABLE "team_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "team_members" ADD FOREIGN KEY ("role") REFERENCES "role" ("name");

ALTER TABLE "todo_actor" ADD FOREIGN KEY ("todo_id") REFERENCES "todo" ("id");

ALTER TABLE "todo_actor" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE INDEX ON "team" ("name");

CREATE INDEX ON "todo" ("name");

CREATE UNIQUE INDEX ON "team_members" ("team_id", "user_id");

CREATE UNIQUE INDEX ON "todo_actor" ("todo_id", "user_id");