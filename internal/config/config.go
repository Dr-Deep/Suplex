package config

type Configuration struct {
	Token string

	Logging struct {
		Method []string
		File   string
		Level  string
	}
}
