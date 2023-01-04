package config

import (
	"github.com/spf13/viper"

	"github.com/cvetkovski98/zvax-common/pkg/config"
)

type Config struct {
	PostgreSQL config.PostgreSQL `mapstructure:"db"`
	MinIO      config.MinIO      `mapstructure:"minio"`
}

func LoadConfig(name string) error {
	viper.AddConfigPath("config")
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func GetConfig() *Config {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}
