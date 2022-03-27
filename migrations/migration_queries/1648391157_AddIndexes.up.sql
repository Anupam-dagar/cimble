CREATE UNIQUE INDEX idx_projectid_name
ON configurations (project_id, name);

CREATE INDEX idx_organisationid
ON projects (organisation_id);