package config

type Config struct {
	DBHost       string
	KafkaBrokers string
	ServicePort  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:       "",      // Load from environment or defaults
		KafkaBrokers: "",      // Load from environment or defaults
		ServicePort:  ":8080", // Default port
	}
}
