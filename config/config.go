package config

type Config struct {
	Port        string
	DatabaseUrl string
}

func NewConfig() *Config {
	return &Config{
		Port:        ":8000",
		DatabaseUrl: "postgres://postgres:Unclesnoopdog@69@localhost:5432/go_clean",
	}
}
