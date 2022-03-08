DROP TABLE IF EXISTS user_organisation_mappings;
DROP TABLE IF EXISTS user_organisation_mapping_archives;

RENAME TABLE user_privileges TO user_mappings;
RENAME TABLE user_privileges_archives TO user_mappings_archives;