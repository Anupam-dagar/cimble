CREATE TABLE user_password (
    user_id varchar(255) NOT NULL,
    password_hash char(128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB;