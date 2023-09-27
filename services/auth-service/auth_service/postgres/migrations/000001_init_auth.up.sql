CREATE TABLE auth_methods (
    "user_id" varchar(64) PRIMARY KEY,
    "password" varchar(128) NOT NULL
);