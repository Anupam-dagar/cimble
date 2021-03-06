package constants

type ErrorMessage string

const (
	Unauthorised     ErrorMessage = "user is unauthorised to perform this action"
	ApiKeyNotPresent ErrorMessage = "missing required query param apiKey"
)
