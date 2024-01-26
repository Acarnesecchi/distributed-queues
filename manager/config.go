package manager

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr string
	ConnMode   string
	MaxNodes   int
	MinNodes   int
	Timeout    time.Duration
}

func NewConfig() Config {
	return Config{
		ListenAddr: "localhost:25255",
		ConnMode:   "tcp",
		MaxNodes:   3,
		MinNodes:   1,
		Timeout:    1 * time.Minute,
	}
}

func NewConfigFromFile(path string) Config {
	v := viper.New()
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config file: %s", path)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("couldn't read config file: %s", path)
	}
	return c
}

func (c Config) WithListenAddr(s string) Config {
	c.ListenAddr = s
	return c
}

func (c Config) WithMaxNodes(n int) Config {
	c.MaxNodes = n
	return c
}

func (c Config) WithMinNodes(n int) Config {
	c.MinNodes = n
	return c
}

func (c Config) WithTimeout(t time.Duration) Config {
	c.Timeout = t
	return c
}

func (c Config) WithConnMode(n string) Config {
	c.ConnMode = n
	return c
}
