CREATE TABLE IF NOT EXISTS periodic_maintenances
(
    id varchar PRIMARY KEY NOT NULL,
    period integer,
    period_unit varchar,
    maintenance_id varchar,
    expires_at timestamp,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (maintenance_id) REFERENCES maintenances (id)
);