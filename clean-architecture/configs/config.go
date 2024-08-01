package configs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type Connections struct {
	Timeout        time.Duration `mapstructure:"timeout"`
	PathConfigFile string        `mapstructure:"path_config_file"`
	Ctx            context.Context
	Paths          *ConfigPath
	FileConfig     *FileConfig
}

type FileConfig struct {
	Extension      string
	Filename       string
	ConfigPath     string
	ConfigFilePath string
}

type ConfigPath struct {
	RootPath       string
	BuildPath      string
	CmdPath        string
	ConfigsPath    string
	InternalPath   string
	PkgPath        string
	ProtoPath      string
	ServicePath    string
	CorePath       string
	EntityPath     string
	RepositoryPath string
	AdapaterPath   string
	HandlerPath    string
	SwaggerPath    string
	InfraPath      string
	StoragePath    string
}

var AppConfig ConfigPath

type Timeout struct{}

func init() {
	_, filename, _, _ := runtime.Caller(0)

	// Root directory
	AppConfig.RootPath = filepath.Dir(filepath.Dir(filename))

	// System Paths
	AppConfig.BuildPath = filepath.Join(AppConfig.RootPath, "build")
	AppConfig.CmdPath = filepath.Join(AppConfig.RootPath, "cmd")
	AppConfig.ConfigsPath = filepath.Join(AppConfig.RootPath, "configs")
	AppConfig.InternalPath = filepath.Join(AppConfig.RootPath, "internal")
	AppConfig.PkgPath = filepath.Join(AppConfig.RootPath, "pkg")
	AppConfig.ProtoPath = filepath.Join(AppConfig.RootPath, "proto")

	// System paths inside internal
	AppConfig.CorePath = (filepath.Join(AppConfig.InternalPath, "core"))
	AppConfig.InfraPath = (filepath.Join(AppConfig.InternalPath, "infra"))

	// System paths inside core
	AppConfig.EntityPath = (filepath.Join(AppConfig.CorePath, "entity"))
	AppConfig.RepositoryPath = (filepath.Join(AppConfig.CorePath, "repository"))
	AppConfig.ServicePath = (filepath.Join(AppConfig.CorePath, "service"))

	// System paths inside infra
	AppConfig.HandlerPath = (filepath.Join(AppConfig.InfraPath, "handler"))
	AppConfig.StoragePath = (filepath.Join(AppConfig.InfraPath, "storage"))
}

func LoadConfig() (*Connections, error) {

	path := os.Getenv("PATH_CONFIG")
	if path == "" {
		path = AppConfig.CmdPath + "/.config"
	}

	fc := FileConfig{
		ConfigPath:     path,
		Extension:      "yaml",
		Filename:       "config",
		ConfigFilePath: fmt.Sprintf("%s/config.yaml", path),
	}

	var cfg *Connections

	viper.AddConfigPath(fc.ConfigPath)
	viper.SetConfigName(fc.Filename)
	viper.SetConfigType(fc.Extension)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err.(viper.ConfigFileNotFoundError)
		}
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	// AppConfig.ConfigFile = AppConfig.ConfigFilePath + "config.yaml"
	err = os.Setenv("JSON_CONFIG_PATH", fc.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	bg := context.Background()

	ctx := context.WithValue(bg, Timeout{}, float64(cfg.Timeout.Seconds()))
	if cfg.Timeout.Seconds() == 0 {
		ctx = context.WithValue(bg, Timeout{}, float64(15))
	}

	i, ok := ctx.Value(Timeout{}).(float64)
	if !ok {
		return nil, errors.New("error to convert timeout to float64")
	}

	return &Connections{
		Timeout:        time.Duration(i) * time.Second,
		PathConfigFile: fc.ConfigFilePath,
		Ctx:            ctx,
		Paths:          &AppConfig,
		FileConfig:     &fc,
	}, err

}
