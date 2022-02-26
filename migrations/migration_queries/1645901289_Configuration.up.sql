CREATE TABLE configuration (
    id varchar(255) NOT NULL DEFAULT "",
    name varchar(255) NOT NULL DEFAULT "",
    info varchar(255) NOT NULL DEFAULT "",
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL DEFAULT "",
    updated_by varchar(255) NOT NULL DEFAULT "",
    PRIMARY KEY (id)
) ENGINE=InnoDB