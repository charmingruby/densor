CREATE TABLE IF NOT EXISTS maintenance_results
(
    id varchar PRIMARY KEY NOT NULL,
    title varchar,
    observation varchar,
    status varchar,
    maintenance_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (maintenance_id) REFERENCES maintenances(id)
);