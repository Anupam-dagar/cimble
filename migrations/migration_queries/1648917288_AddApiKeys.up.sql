CREATE TABLE api_keys (
    id varchar(255) NOT NULL,
    organisation_id varchar(255) NOT NULL,
    key_hash binary(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    revoked TINYINT(1) NOT NULL DEFAULT 0,
    privileges varchar(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB;