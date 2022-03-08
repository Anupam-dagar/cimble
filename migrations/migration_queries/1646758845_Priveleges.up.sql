CREATE TABLE privileges (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(20) NOT NULL,
    is_read TINYINT(1) NOT NULL DEFAULT 1,
    is_write TINYINT(1) NOT NULL DEFAULT 0,
    is_update TINYINT(1) NOT NULL DEFAULT 0,
    is_delete TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO privileges (name, is_read, is_write, is_update, is_delete)
VALUES ("owner", 1, 1, 1, 1), ("admin", 1, 1, 1, 0);