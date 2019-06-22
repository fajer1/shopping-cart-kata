package webapi

// Config represents the app configuration
type Config struct {
	CompanyName   string
	HashSalt      string
	ListenAddress string
	EnableCache   bool
}
