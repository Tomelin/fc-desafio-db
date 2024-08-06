package cache

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type DesafioCache *redis.Client

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

func NewCacheConnection(ctx context.Context, pathConfigFile, nameConfigFile, nameFileExtension string) (*redis.Client, error) {

	viper.AddConfigPath(pathConfigFile)
	viper.SetConfigName(nameConfigFile)
	viper.SetConfigType(nameFileExtension)
	viper.AutomaticEnv()

	v, ok := viper.Get("cache").(map[string]interface{})
	if !ok {
		return nil, errors.New("error, not found cache key at config file")
	}

	cacheConfig := RedisConfig{}
	err := cacheConfig.ValidateFiledFromViper(v)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheConfig.Host, cacheConfig.Port),
		Password: cacheConfig.Password,
		DB:       cacheConfig.Database,
	})

	return rdb, nil
}

func (cacheConfig *RedisConfig) ValidateFiledFromViper(field map[string]interface{}) error {

	cacheConfig.Host = "localhost"
	cacheConfig.Port = "6379"
	cacheConfig.Database = 0

	fInt, ok := field["database"].(int)
	if ok {
		cacheConfig.Database = fInt
	}

	fStr, ok := field["port"].(string)
	if ok {
		cacheConfig.Port = fStr
	}

	fStr, ok = field["host"].(string)
	if ok {
		cacheConfig.Host = fStr
	}

	fStr, ok = field["password"].(string)
	if ok {
		cacheConfig.Password = fStr
	} else {
		return errors.New("password cannot be empty")
	}

	return nil
}
