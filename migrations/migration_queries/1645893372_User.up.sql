CREATE TABLE user (
    id varchar(255) NOT NULL,
    organisation_id varchar(255) NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    email varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(email)
) ENGINE=InnoDB;