CREATE TABLE IF NOT EXISTS users 
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    email varchar UNIQUE,
    password varchar,
    role varchar,
    is_active boolean,
    created_at timestamp DEFAULT now() NOT NULL
);