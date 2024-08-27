CREATE TABLE IF NOT EXISTS sensor_failures
(
    id varchar PRIMARY KEY NOT NULL,
    title varchar, 
    description varchar,
    rate float,
    sensor_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,


FOREIGN KEY (sensor_id) REFERENCES sensors(id) );