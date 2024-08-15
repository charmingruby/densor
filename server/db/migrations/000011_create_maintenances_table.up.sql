CREATE TABLE IF NOT EXISTS maintenances
(
    id varchar PRIMARY KEY NOT NULL,
    title varchar,
    start_date timestamp,
    reason varchar,
    observation varchar,
    description varchar,
    equipment_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (equipment_id) REFERENCES equipments (id)
);