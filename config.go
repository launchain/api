package api

import "regexp"

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
func (c *Config) Check() {
	re := regexp.MustCompile(`^((?:2[0-5]{2}|1\d{2}|[1-9]\d|[1-9])\.(?:(?:2[0-5]{2}|1\d{2}|[1-9]\d|\d)\.){2}(?:2[0-5]{2}|1\d{2}|[1-9]\d|\d)):(\d|[1-9]\d|[1-9]\d{2,3}|[1-5]\d{4}|6[0-4]\d{3}|654\d{2}|655[0-2]\d|6553[0-5])$`)
	uri := c.Host + ":" + c.Port
	match := re.MatchString(uri)
	if !match {
		re1 := regexp.MustCompile(`^([a-zA-Z\d][a-zA-Z\d-_]+\.)+[a-zA-Z\d-_][^ ]*:(\d|[1-9]\d|[1-9]\d{2,3}|[1-5]\d{4}|6[0-4]\d{3}|654\d{2}|655[0-2]\d|6553[0-5])$`)
		match = re1.MatchString(uri)
	}
	if !match {
		panic("uri must like \"192.168.200.132:5006\"0 or \"baidu.com:80\"")
	}
}

// URI ...
func (c *Config) URI() string {
	return "http://" + c.Host + ":" + c.Port
}
