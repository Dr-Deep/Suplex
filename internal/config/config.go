package config

type Configuration struct {
	Prefix string `json:"prefix"`
	Token  string `json:"token"`

	Logging struct {
		Method []string `json:"method"`
		File   string   `json:"file"`
		Level  string   `json:"level"`
	} `json:"logging"`
}
