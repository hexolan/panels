CREATE TABLE posts (
    "id" serial PRIMARY KEY,
    "panel_id" varchar(64) NOT NULL,
    "author_id" varchar(64) NOT NULL,
    "title" varchar(512) NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT timezone('utc', now()),
    "updated_at" timestamp
);