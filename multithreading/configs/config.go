package configs

//
//import (
//	"fmt"
//	"github.com/spf13/viper"
//)
//
//type ConfigWebserver struct {
//	Port         string `json:"port" mapstructure:"port"`
//	EnabledHttp2 bool   `json:"disable_http2" mapstructure:"enabled_http2"`
//	Listen       string `json:"listen" mapstructure:"listen"`
//}
//
//type ConfigDatabase struct {
//	Driver   string `json:"driver" mapstructure:"driver"`
//	Host     string `json:"host" mapstructure:"host"`
//	Username string `json:"username" mapstructure:"username"`
//	Password string `json:"password" mapstructure:"password"`
//}
//
//type Config struct {
//	Database  *ConfigDatabase  `json:"database" mapstructure:"database"`
//	Webserver *ConfigWebserver `json:"webserver" mapstructure:"webserver"`
//}
//
//func NewConfig(f string) (*Config, error) {
//	viper.AddConfigPath(f)
//	viper.SetConfigName("config") // Register config file name (no extension)
//	viper.SetConfigType("yaml")   // Look for specific type
//	err := viper.ReadInConfig()
//	if err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
//			panic(fmt.Errorf("error: %s", err.Error()))
//		} else {
//			panic(fmt.Errorf("fatal error config file: %w", err))
//		}
//
//	}
//	c := &Config{}
//	err = viper.Unmarshal(c)
//	if err != nil {
//		panic(fmt.Errorf("fatal error to make struct config: %w", err))
//	}
//
//	return c, nil
//}
