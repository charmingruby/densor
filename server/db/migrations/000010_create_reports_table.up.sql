CREATE TABLE IF NOT EXISTS reports
(
    id varchar PRIMARY KEY NOT NULL,
    title varchar,
    type varchar,
    report_url varchar,
    sector_id varchar,
    author_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,


FOREIGN KEY (sector_id) REFERENCES sectors (id),
FOREIGN KEY (author_id) REFERENCES users (id)

);