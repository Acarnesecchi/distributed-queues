package worker

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	TargetAddr string
	ConnMode   string
	Tasks      []string
	Timeout    time.Duration
}

func NewConfig() Config {
	return Config{
		TargetAddr: "localhost:25255",
		ConnMode:   "tcp",
		Timeout:    5 * time.Second,
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

func (c Config) WithTimeout(t time.Duration) Config {
	c.Timeout = t
	return c
}

func (c Config) WithConnMode(n string) Config {
	c.ConnMode = n
	return c
}

func (c Config) WithTargetAddr(s string) Config {
	c.TargetAddr = s
	return c
}

func (c Config) WithTasks(t ...string) Config {
	c.Tasks = t
	return c
}
