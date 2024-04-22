package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigWebserver struct {
	Port         string `json:"port" mapstructure:"port"`
	EnabledHttp2 bool   `json:"disable_http2" mapstructure:"enabled_http2"`
	Listen       string `json:"listen" mapstructure:"listen"`
}

type ConfigDatabase struct {
	Driver   string `json:"driver" mapstructure:"driver"`
	Host     string `json:"host" mapstructure:"host"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

type Config struct {
	Database  *ConfigDatabase  `json:"database" mapstructure:"database"`
	Webserver *ConfigWebserver `json:"webserver" mapstructure:"webserver"`
}

func NewConfig(f string) (*Config, error) {
	viper.AddConfigPath(f)
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("error: %s \n FILE_CONFIG environment must be setted to load .config/config.yaml file", err.Error())
		} else {
			return nil, fmt.Errorf("fatal error config file: %s  \n FILE_CONFIG environment must be setted to load .config/config.yaml file", err.Error())
		}

	}
	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("fatal error to make struct config: %w", err)
	}

	return c, nil
}
