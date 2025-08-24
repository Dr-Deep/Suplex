package config

type DiscordConfig struct {
	Token          string
	Application_ID string
	Client_ID      string
	Client_Secret  string
	PublicKey      string
	Guild_ID       string
}

type Configuration struct {
	Discord_Settings DiscordConfig
}
