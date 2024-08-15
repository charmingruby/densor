CREATE TABLE IF NOT EXISTS equipments 
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    model varchar,
    manufacturer varchar,
    sector_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (sector_id) REFERENCES sectors (id) ON DELETE SET NULL
);