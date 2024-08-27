CREATE TABLE IF NOT EXISTS alarm
(
    id varchar PRIMARY KEY NOT NULL,
    type varchar,
    description varchar,
    occurred_at timestamp,
    failure_notification_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (failure_notification_id) REFERENCES failure_notifications(id)
);