CREATE TABLE panels (
    "id" serial PRIMARY KEY,
    "name" varchar(32) NOT NULL,
    "description" varchar(512) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT timezone('utc', now()),
    "updated_at" timestamp
);
CREATE UNIQUE INDEX panels_name_unique ON "panels" (LOWER("name"));

INSERT INTO panels ("name", "description") VALUES ('Panel', 'The de facto panel.') ON CONFLICT DO NOTHING;