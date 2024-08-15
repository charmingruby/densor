CREATE TABLE IF NOT EXISTS failure_notifications
(
    id varchar PRIMARY KEY NOT NULL,
    message varchar,
    occurred_at timestamp,
    sector_id varchar,
    sensor_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,


    FOREIGN KEY (sector_id) REFERENCES sectors(id),
    FOREIGN KEY (sensor_id) REFERENCES sensors(id) 
);