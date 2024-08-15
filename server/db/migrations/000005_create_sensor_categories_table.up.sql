CREATE TABLE IF NOT EXISTS sensor_categories (
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    description varchar,
    created_at timestamp DEFAULT now() NOT NULL
);