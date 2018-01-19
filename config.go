package api

// Config ...
type Config struct {
	Host string
	Port string
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: "8888",
	}
}

// Check ...
func (c *Config) Check() bool {
	// TODO
	return true
}
