CREATE TABLE "user"
(
    "id"       serial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "password" varchar        NOT NULL
);

CREATE TABLE "note"
(
    "id"        SERIAL PRIMARY KEY,
    "note"      text,
    "author_id" int NOT NULL
);

ALTER TABLE "note"
    ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");