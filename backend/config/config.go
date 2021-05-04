package config

type Config struct {
	ServerConfig `yaml:"server"`
	MongoConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type MongoConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}