package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

type PostgresSqlConfig struct {
	Host                  string `mapstructure:"host"`
	Username              string `mapstructure:"username"`
	Port                  string `mapstructure:"port"`
	Database              string `mapstructure:"database"`
	Password              string `mapstructure:"password"`
	SSLEnabled            string `mapstructure:"ssl_enabled"`
	MaxConnection         int32  `mapstructure:"max_connection"`
	MaxIdleConnection     int32  `mapstructure:"max_idle_connection"`
	MinConnection         int32  `mapstructure:"min_connection"`
	MaxConnectionLifeTime int32  `mapstructure:"max_connection_lifetime"`
	MaxConnectionIdleTime int32  `mapstructure:"max_connection_idletime"`
	Options               string `mapstructure:"options"`
	Offset                int8   `mapstructure:"offset"`
	Limit                 int8   `mapstructure:"limit"`
}

func NewDBConnection(ctx context.Context, pathConfigFile, nameConfigFile, nameFileExtension string) (*sql.DB, error) {

	viper.AddConfigPath(pathConfigFile)
	viper.SetConfigName(nameConfigFile)
	viper.SetConfigType(nameFileExtension)
	viper.AutomaticEnv()

	v, ok := viper.Get("postgresql").(map[string]interface{})
	if !ok {
		return nil, errors.New("error, not found potsgres key at config file")
	}

	sqlConfig := PostgresSqlConfig{}
	err := sqlConfig.ValidateFiledFromViper(v)
	if err != nil {
		return nil, err
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", sqlConfig.Username, sqlConfig.Password, sqlConfig.Host, sqlConfig.Port, sqlConfig.Database)
	pool, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(time.Duration(sqlConfig.MaxConnectionLifeTime))
	pool.SetMaxIdleConns(int(sqlConfig.MaxIdleConnection))
	pool.SetMaxOpenConns(int(sqlConfig.MaxConnection))

	// ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	// if err := pool.PingContext(ctxPing); err != nil {
	// 	return nil, err
	// }

	return pool, nil
}

func (sqlConfig *PostgresSqlConfig) ValidateFiledFromViper(field map[string]interface{}) error {

	sqlConfig.Host = "localhost"
	sqlConfig.Port = "5432"
	sqlConfig.SSLEnabled = "false"
	sqlConfig.Password = ""
	sqlConfig.MaxConnection = int32(5)
	sqlConfig.MinConnection = 1
	sqlConfig.MaxConnectionLifeTime = 10
	sqlConfig.MaxConnectionIdleTime = 5
	sqlConfig.MaxIdleConnection = 3
	sqlConfig.Limit = 100
	sqlConfig.Offset = 0

	fInt32, ok := field["max_connection"].(int32)
	if ok {
		sqlConfig.MaxConnection = fInt32
	}

	fInt32, ok = field["min_connection"].(int32)
	if ok {
		sqlConfig.MinConnection = fInt32
	}

	fInt32, ok = field["max_connection_lifetime"].(int32)
	if ok {
		sqlConfig.MaxConnectionLifeTime = fInt32
	}

	fInt32, ok = field["max_connection_idletime"].(int32)
	if ok {
		sqlConfig.MaxConnectionIdleTime = fInt32
	}

	fInt32, ok = field["max_idle_connection"].(int32)
	if ok {
		sqlConfig.MaxIdleConnection = fInt32
	}

	fStr, ok := field["port"].(string)
	if ok {
		sqlConfig.Port = fStr
	}

	fStr, ok = field["ssl_enabled"].(string)
	if ok {
		sqlConfig.SSLEnabled = fStr
	}

	fStr, ok = field["host"].(string)
	if ok {
		sqlConfig.Host = fStr
	}

	fStr, ok = field["options"].(string)
	if ok {
		sqlConfig.Options = fStr
	}

	fint8, ok := field["limit"].(int8)
	if ok {
		sqlConfig.Limit = fint8
	}

	fint8, ok = field["offset"].(int8)
	if ok {
		sqlConfig.Offset = fint8
	}

	fStr, ok = field["database"].(string)
	if ok {
		sqlConfig.Database = fStr
	} else {
		return errors.New("database cannot be empty")
	}

	fStr, ok = field["username"].(string)
	if ok {
		sqlConfig.Username = fStr
	} else {
		return errors.New("username cannot be empty")
	}

	fStr, ok = field["password"].(string)
	if ok {
		sqlConfig.Password = fStr
	} else {
		return errors.New("password cannot be empty")
	}

	return nil
}
