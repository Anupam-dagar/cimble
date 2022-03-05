CREATE TABLE configurations (
    id varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    info varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB