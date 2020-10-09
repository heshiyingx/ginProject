package config

import (
	"database/sql"
	"time"
)

// Server 服务配置
type Server struct {
	ID       int64  `yaml:"id"`
	Debug    bool   `yaml:"debug"`
	Port     string `yaml:"port"`
	Timezone string `yaml:"timezone"`
	Location *time.Location
}

// MysqlConf ...
type MysqlConf struct {
	UserName string `yaml:"username"`
	PWD      string `yaml:"pwd"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// RedisConf ...
type RedisConf struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"pwd"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"poolSize"`
	DialTimeout  int    `yaml:"dialTimeout"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}

// Config 配置写这里，传递到各个package
type Config struct {
	FileName      string
	Authorization bool      `yaml:"authorization"`
	Server        Server    `yaml:"server"`
	GoodsDB       MysqlConf `yaml:"mysql"`
	RedisConf     RedisConf `yaml:"redis"`
	DB            *sql.DB   `yaml:"db"`
	Env           string    `yaml:"env"`
}
