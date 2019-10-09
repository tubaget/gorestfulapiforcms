package setting

import (
	"sync"

	"github.com/BurntSushi/toml"

	"path/filepath"

	"fmt"
	"time"
)

var config_path = "config/config.toml"

type toml_config struct {
	Run_mode string
	Database *database
	Server   *server
	Basic    *basic
	Log      *log
}

type database struct {
	Type         string
	Host         string
	Port         int
	User         string
	Password     string
	Db_name      string `toml:"database_name"`
	Table_prefix string
}

type server struct {
	Http_port     int
	Read_timeout  time.Duration
	Write_timeout time.Duration
}

type basic struct {
	Page_size int
	Jwt_secret string
}

type log struct {
	Log_path string
}

var (
	cfg  *toml_config
	once sync.Once
)

func Config() *toml_config {
	once.Do(func() {
		filePath, err := filepath.Abs(config_path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
