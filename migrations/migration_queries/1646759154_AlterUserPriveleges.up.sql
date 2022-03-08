ALTER TABLE user_privileges
DROP COLUMN is_read,
DROP COLUMN is_write,
DROP COLUMN is_update,
DROP COLUMN is_delete;

ALTER TABLE user_privileges
ADD privelege varchar(20) NOT NULL;

ALTER TABLE user_privileges_archives
DROP COLUMN is_read,
DROP COLUMN is_write,
DROP COLUMN is_update,
DROP COLUMN is_delete;

ALTER TABLE user_privileges_archives
ADD privelege varchar(20) NOT NULL;