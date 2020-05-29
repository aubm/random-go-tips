package config

type Config struct {
	WebAppAddr string
}

func NewWithDefaults() Config {
	return Config{
		WebAppAddr: ":8080",
	}
}
