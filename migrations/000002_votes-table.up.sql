CREATE TABLE IF NOT EXISTS votes (
    id                          UUID NOT NULL PRIMARY KEY,
    voter_id                    UUID NOT NULL,
    poll_id                     UUID NOT NULL,
    choosen_options             VARCHAR(255)[] NOT NULL,
    created_at                  DATE NOT NULL,
    FOREIGN KEY(poll_id)        REFERENCES polls(id),
    CONSTRAINT uc_poll_voter    UNIQUE (poll_id, voter_id)
);