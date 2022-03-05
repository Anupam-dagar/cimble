CREATE TABLE user_privileges_archive (
    user_id varchar(255) NOT NULL,
    level_for ENUM('organisation', 'project') NOT NULL DEFAULT "project",
    level_id varchar(255) NOT NULL,
    is_read TINYINT(1) NOT NULL DEFAULT 1,
    is_write TINYINT(1) NOT NULL DEFAULT 0,
    is_update TINYINT(1) NOT NULL DEFAULT 0,
    is_delete TINYINT(1) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    deleted_timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_by varchar(255) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB