package steps

type Guide int

const (
	Welcome Guide = iota
	CreateDatabase
	GetDatabaseID
	CreateIntegration
	ConnectDatabase
	GetToken
	TestConnection
	Complete
)
