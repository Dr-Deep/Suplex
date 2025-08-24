package config

type DiscordConfig struct {
	Token          string `yaml:"Token"`
	Application_ID string `yaml:"Application_ID"`
	Client_ID      string `yaml:"Client_ID"`
	Client_Secret  string `yaml:"Client_Secret"`
	PublicKey      string `yaml:"PublicKey"`
	Guild_ID       string `yaml:"Guild_ID"`
}

type Configuration struct {
	Discord_Settings DiscordConfig `yaml:"discord"`
	AssetsDir        string        `yaml:"Assets_Dir"`
}
