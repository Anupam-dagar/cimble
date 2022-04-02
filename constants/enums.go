package constants

type Privileges string

const (
	OWNER Privileges = "owner"
	ADMIN Privileges = "admin"
)

type PrivilegeLevel string

const (
	ORGANISATION PrivilegeLevel = "organisation"
	PROJECT      PrivilegeLevel = "project"
)

type SecurityType string

const (
	PASSWORD SecurityType = "password"
	ApiKey   SecurityType = "api_key"
)
