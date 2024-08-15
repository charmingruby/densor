CREATE TABLE IF NOT EXISTS sectors
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar,
    localization varchar,
    responsible_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE SET NULL
);

