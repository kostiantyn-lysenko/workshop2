package storage

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost dbname=postgres sslmode=disable user=postgres password=password",
	}
}
