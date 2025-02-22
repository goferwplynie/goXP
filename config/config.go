package config

type Config struct {
	Cursor                string
	MaxFilesToLoadAtStart int
}

func New() *Config {
	return &Config{
		Cursor:                ">",
		MaxFilesToLoadAtStart: 100,
	}
}
