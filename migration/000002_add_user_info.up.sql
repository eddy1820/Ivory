CREATE TABLE user_info
(
    id         bigint unsigned NOT NULL AUTO_INCREMENT,
    account_id bigint unsigned NOT NULL,
    gender     varchar(255) NOT NULL DEFAULT '',
    name       varchar(255) NOT NULL DEFAULT '',
    address    varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES account_info (id),
    UNIQUE KEY idx_account_id (account_id)
);
