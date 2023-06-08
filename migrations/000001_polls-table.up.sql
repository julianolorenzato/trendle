CREATE TABLE IF NOT EXISTS polls (
    id                  UUID NOT NULL PRIMARY KEY,
    question            VARCHAR(255) NOT NULL,
    number_of_choices   INTEGER NOT NULL,
    options             VARCHAR(255)[] NOT NULL,
    is_permanent        BOOLEAN NOT NULL,
    expires_at          DATE NOT NULL,
    created_at          DATE NOT NULL
);