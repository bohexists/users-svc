package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server      ServerConfig      `yaml:"server"`
	RateLimiter RateLimiterConfig `yaml:"rate_limiter"`
	Cache       CacheConfig       `yaml:"cache"`
	Mongo       MongoConfig       `yaml:"mongo"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type RateLimiterConfig struct {
	MaxRequestsPerSecond int `yaml:"max_requests_per_second"`
	BurstSize            int `yaml:"burst_size"`
}

type CacheConfig struct {
	MaxSize    int    `yaml:"max_size"`
	DefaultTTL string `yaml:"default_ttl"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type MongoConfig struct {
	URI        string `yaml:"uri"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}
