CREATE TABLE IF NOT EXISTS stock_materials
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    stock_quantity integer,
    minimum_quantity int,
    created_at timestamp DEFAULT now() NOT NULL
);