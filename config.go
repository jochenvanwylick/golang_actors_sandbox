package main

type Config struct {
	Setting string
}

func NewConfig() *Config {
	return &Config{
		Setting: "test",
	}
}
