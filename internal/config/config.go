package config

type Config struct {
	DBPath string
}

func Load() *Config {
	return &Config{
		DBPath: "test.db", // For simplicity, using SQLite. In production, use environment variables.
	}
}
