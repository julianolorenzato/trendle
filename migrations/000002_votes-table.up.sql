CREATE TABLE IF NOT EXISTS votes (
    id                  UUID NOT NULL PRIMARY KEY,
    voter_id            UUID NOT NULL,
    options_choosed     VARCHAR(255)[] NOT NULL,
    created_at          DATE NOT NULL,
    poll_id             UUID NOT NULL,
    FOREIGN KEY(poll_id) REFERENCES polls(id)
);