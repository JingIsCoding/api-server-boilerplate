package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var config *Config
var once sync.Once

type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
	LogConfig      LogConfig      `mapstructure:"log"`
}

type DatabaseConfig struct {
	Type        string `mapstructure:"type"`
	DatabaseUrl string `mapstructure:"url"`
}

type ServerConfig struct {
	ServerName string `mapstructure:"service_name"`
	HMACSecret string `mapstructure:"hmac_secret"`
	Domain     string `mapstructure:"domain"`
	Port       string `mapstructure:"port"`
}

type RedisConfig struct {
	URL         string `mapstructure:"url"`
	Name        string `mapstructure:"name"`
	Group       string `mapstructure:"group"`
	Host        string `mapstructure:"host"`
	Port        int64  `mapstructure:"port"`
	Password    string `mapstructure:"password"`
	MaxIdle     int64  `mapstructure:"maxidle"`
	MaxActive   int64  `mapstructure:"maxactive"`
	IdleTimeout int64  `mapstructure:"idletimeout"`
	IsMaster    bool   `mapstructure:"ismaster"`
	Wait        bool   `mapstructure:"wait"`
}

type LogConfig struct {
	LogLevel string `mapstructure:"log_level"`
}

func Get() *Config {
	once.Do(func() {
		var conf Config
		v := viper.New()
		loadFromEnv(v)
		if err := v.Unmarshal(&conf); err != nil {
			log.Fatal(err)
		}
		config = &conf
	})
	return config
}

func loadFromEnv(v *viper.Viper) {
	// SERVER
	v.BindEnv("server.service_name", "SERVER_SERVICE_NAME")
	v.BindEnv("server.hmac_secret", "SERVER_HMAC_SECRET")
	v.BindEnv("server.domain", "SERVER_DOMAIN")
	v.BindEnv("server.port", "SERVER_PORT")

	// REDIS_URL
	v.BindEnv("redis.url", "REDIS_URL")
	v.BindEnv("redis.maxidle", "REDIS_MAX_IDLE")
	v.BindEnv("redis.maxactive", "REDIS_MAX_ACTIVE")
	v.BindEnv("redis.wait", "REDIS_WAIT")
	v.BindEnv("redis.idletimeout", "REDIS_IDLETIMEOUT")

	// DataBase
	v.BindEnv("database.url", "DATABASE_URL")
	v.BindEnv("database.type", "DATABASE_TYPE")

	// Log
	v.BindEnv("log.log_level", "LOG_LEVEL")
}
