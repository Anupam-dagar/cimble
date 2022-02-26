CREATE TABLE user_privileges_archive (
    id varchar(255) NOT NULL DEFAULT "",
    user_id varchar(255) NOT NULL DEFAULT "",
    level_for ENUM('organisation', 'project') NOT NULL DEFAULT "project",
    level_id varchar(255) NOT NULL DEFAULT "",
    is_read TINYINT(1) NOT NULL DEFAULT 1,
    is_write TINYINT(1) NOT NULL DEFAULT 0,
    is_update TINYINT(1) NOT NULL DEFAULT 0,
    is_delete TINYINT(1) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL DEFAULT "",
    updated_by varchar(255) NOT NULL DEFAULT "",
    deleted_timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_by varchar(255) NOT NULL DEFAULT "",
    PRIMARY KEY (id)
) ENGINE=InnoDB