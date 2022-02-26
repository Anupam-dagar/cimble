CREATE TABLE organisation_archive (
    id varchar(255) NOT NULL DEFAULT "",
    name varchar(255) NOT NULL DEFAULT "",
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL DEFAULT "",
    updated_by varchar(255) NOT NULL DEFAULT "",
    deleted_timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_by varchar(255) NOT NULL DEFAULT "",
    PRIMARY KEY (id)
) ENGINE=InnoDB