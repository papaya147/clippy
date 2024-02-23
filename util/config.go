package util

type Config struct {
	StoreFile string `json:"store_file"`
}

func NewConfig() *Config {
	return &Config{
		StoreFile: "~/clippy/store.json",
	}
}

func DefaultConfig() *Config {
	return NewConfig()
}
