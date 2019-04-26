-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    realm_id    VARCHAR(32) NOT NULL,
    user_id     VARCHAR(32) NOT NULL,
    principal   TEXT NOT NULL,
    secret      TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    PRIMARY KEY(realm_id, user_id),
    CONSTRAINT unq_user_principal UNIQUE (realm_id, principal)
);

-- +migrate Down
DROP TABLE users;