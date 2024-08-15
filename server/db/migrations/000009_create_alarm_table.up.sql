CREATE TABLE IF NOT EXISTS alarm
(
    id varchar PRIMARY KEY NOT NULL,
    type varchar,
    description varchar,
    occurred_at timestamp,
    sector_id varchar,
    sensor_id varchar,
    failure_notification_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,


    FOREIGN KEY (sector_id) REFERENCES sectors(id),
    FOREIGN KEY (sensor_id) REFERENCES sensors(id),
    FOREIGN KEY (failure_notification_id) REFERENCES failure_notifications(id)
);