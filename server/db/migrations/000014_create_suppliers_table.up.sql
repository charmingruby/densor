CREATE TABLE IF NOT EXISTS suppliers
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    area_of_activity varchar,
    responsible_name varchar,
    email varchar,
    phone_number varchar,
    address varchar,
    created_at timestamp DEFAULT now() NOT NULL
);