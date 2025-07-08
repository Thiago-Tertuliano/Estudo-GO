package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	API  APIConfig `mapstructure:"api"`
	Database DBConfig `mapstructure:"database"`
}

type APIConfig struct {
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal da configuração: %w", err)
	}

	return &config, nil
}

func (c *Config) GetDBConfig() DBConfig {
	return c.Database
}

func (c *Config) GetAPIConfig() APIConfig {
	return c.API
}