CREATE TABLE IF NOT EXISTS occasional_maintenances
(
    id varchar PRIMARY KEY NOT NULL,
    maintenance_id varchar,
    deadline_at timestamp,
    concluded_at timestamp,
    created_at timestamp DEFAULT now() NOT NULL,
    
    FOREIGN KEY (maintenance_id) REFERENCES maintenances (id)
);