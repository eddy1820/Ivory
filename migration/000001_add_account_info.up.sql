CREATE TABLE IvoryDb.account
(
    id                  bigint unsigned NOT NULL AUTO_INCREMENT,
    account             varchar(255) NOT NULL DEFAULT '',
    hashed_password     varchar(255) NOT NULL DEFAULT '',
    email               varchar(255) NOT NULL DEFAULT '',
    password_changed_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY idx_account (account),
    UNIQUE KEY idx_email (email)
);
