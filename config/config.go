package config

type Config struct {
	Cursor string
}

func New() *Config {
	return &Config{
		Cursor: ">",
	}
}
