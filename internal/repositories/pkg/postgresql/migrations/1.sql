-- +migrate Up
CREATE TABLE IF NOT EXISTS realms (
    realm_id    VARCHAR(32) NOT NULL PRIMARY KEY,
    label       VARCHAR(50) NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    CONSTRAINT unq_realm_label UNIQUE (label)
);

-- +migrate Down
DROP TABLE realms;