ALTER TABLE user_privileges DROP PRIMARY KEY;
ALTER TABLE user_privileges ADD PRIMARY KEY (user_id, level_id);

ALTER TABLE user_privileges_archives DROP PRIMARY KEY;
ALTER TABLE user_privileges_archives ADD PRIMARY KEY (user_id, level_id);