package common

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v2"
)

// global var
var (
	G_AppConfig AppConfig
	G_DBConfig  DBConfig
)

// 全局配置文件对应的结构体
type (
	// app
	AppConfig struct {
		iris.Configuration `yaml:"Configuration"`
		Own                `yaml: "own"`
	}
	Own struct {
		Separate      bool     `yaml:"separate"` // 是否前后端分离
		Port          int      `yaml:"port"`
		IgnoreURLs    []string `yaml:"ignore_urls,flow"`
		InterceptURLs []string `yaml:"intercept_urls,flow"`
		JWTTimeout    int      `yaml:"jwt_timeout"`
		LogLevel      string   `yaml:"log_level"`
		Secret        string   `yaml:"secret"`
		WebsocketPool int      `yaml:"websocket_pool"`
		Domain        string   `yaml:"domain"`
	}

	// db
	DBConfig struct {
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
			PoolSize int    `yaml:"poolSize"`
		}
		Mongodb struct {
			Addr           string `yaml:"addr"`
			Database       string `yaml:"database"`
			ConnectTimeout int    `yaml:"connectTimeout"`
			MaxPoolSize    int    `yaml:"maxPoolSize"`
		}
		Mysql struct {
			Dialect      string `yaml:"dialect"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			Database     string `yaml:"database"`
			Charset      string `yaml:"charset"`
			ShowSql      bool   `yaml:"showSql"`
			LogLevel     string `yaml:"logLevel"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
			//ParseTime       bool   `yaml:"parseTime"`
			//MaxIdleConns    int    `yaml:"maxIdleConns"`
			//MaxOpenConns    int    `yaml:"maxOpenConns"`
			//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
			//Sslmode         string `yaml:"sslmode"`
		}
	}
)

func (conf DBConfig) DBConnUrl() string {
	var info = conf.Mysql
	//"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
	//return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
}

func InitConfig(asset func(name string) ([]byte, error)) {
	var (
		app  = AppConfig{}
		db   = DBConfig{}
		data []byte
		err  error
	)
	// app
	data, err = asset("conf/app.yaml")
	if err != nil {
		goto ERR
	}
	if err = yaml.Unmarshal(data, &app); err != nil {
		goto ERR
	}
	G_AppConfig = app
	golog.Infof("[app config]=> %v", G_AppConfig)

	// db
	data, err = asset("conf/db.yaml")
	if err != nil {
		goto ERR
	}
	if err = yaml.Unmarshal(data, &db); err != nil {
		goto ERR
	}
	G_DBConfig = db
	golog.Infof("[db  config]=> %v", G_DBConfig)
	return
ERR:
	golog.Fatalf("~~> 解析配置文件错误,原因:%s", err.Error())
}
