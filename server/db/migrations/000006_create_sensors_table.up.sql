CREATE TABLE IF NOT EXISTS sensors 
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    description varchar,
    sensor_category_id varchar,
    equipment_id varchar,
    observation varchar,
    sector_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,


    FOREIGN KEY (sensor_category_id) REFERENCES sensor_categories(id),
    FOREIGN KEY (sector_id) REFERENCES sectors(id),
    FOREIGN KEY (equipment_id) REFERENCES equipments(id)
);