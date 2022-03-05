CREATE TABLE user_organisation_mappings (
    id varchar(255) NOT NULL,
    user_id varchar(255) NOT NULL,
    organisation_id varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB