CREATE TABLE comments (
    "id" serial PRIMARY KEY,
    
    "post_id" varchar(64) NOT NULL,
    "author_id" varchar(64) NOT NULL,

    "message" varchar(512) NOT NULL,

    "created_at" timestamp NOT NULL DEFAULT timezone('utc', now()),
    "updated_at" timestamp
);