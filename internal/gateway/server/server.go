package server

import (
	"context"
	"fmt"
	"gatewayH/internal/gateway/config"
	"gatewayH/internal/gateway/services"
	"gatewayH/pkg/idgenerator"
	"gatewayH/utils"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run is used to run server.
func Run(filename string) {

	var config config.Config
	config.FileName = filename

	s, err := InitServer(&config)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	srv := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: initRouter(s),
	}

	go func() {
		// service connections
		log.Println("Run server successfully", s.Cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -3 is syscall.SIGQUIT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

func InitServer(config *config.Config) (*services.Services, error) {
	var err error
	// 初始化配置
	if err = InitConfigFromFile(config); err != nil {
		return nil, err
	}

	// server Logger
	logger := InitLogger(config.Server.Debug)
	logger.ReportCaller = true
	// 初始化数据库

	mainConn, err := InitMysql(config, &config.GoodsDB)
	if err != nil {
		return nil, err
	}

	// 初始化redis
	redis, err := InitRedis(config)
	if err != nil {
		return nil, err
	}

	idWorker, err := idgenerator.NewIdWorker(config.Server.ID)
	if err != nil {
		panic(err.Error())
	}

	s := services.NewServices(config, mainConn, logger, redis, idWorker)

	return s, nil
}

// InitLogger 初始化日志系统
func InitLogger(debug bool) *logrus.Logger {

	logger := log.New()

	// 设置日志格式为json格式
	logger.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logger.SetOutput(os.Stdout)

	// 正式发版设置日志级别为warn以上
	if !debug {
		logger.SetLevel(log.WarnLevel)
	} else {
		logger.SetLevel(log.TraceLevel)
		log.Println("\x1b[32mDEBUG MODE\x1b[0m")
	}
	return logger

}

// InitConfigFromFile ...
func InitConfigFromFile(config *config.Config) error {

	// 1.解析 YAML 配置文件
	bytes, err := ioutil.ReadFile(config.FileName)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return err
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	if err != nil {
		log.Fatalf("LoadLocation error: %v", err)
		return err
	}
	return nil
}

// InitMysql ...
func InitMysql(config *config.Config, mysqlConfig *config.MysqlConf) (*dbr.Connection, error) {
	var (
		USERNAME = mysqlConfig.UserName
		PASSWORD = mysqlConfig.PWD
		NETWORK  = mysqlConfig.Network
		SERVER   = mysqlConfig.Server
		PORT     = mysqlConfig.Port
		DATABASE = mysqlConfig.Database
	)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	logger := utils.NewDBLogger(config.Server.Debug)
	conn, err := dbr.Open("mysql", dsn, logger)
	if err != nil {
		log.Printf("Open mysql failed, err:%v\n", err)
		return nil, err
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(0)
	if err := conn.Ping(); err != nil {
		conn.Close()
		log.Printf("Connection mysql ping failed, err:%v\n", err)
		return nil, err
	}

	return conn, nil
}

// InitRedis 初始化redis client
func InitRedis(config *config.Config) (r *redis.Client, err error) {
	opts := redis.Options{
		Addr:         config.RedisConf.Addr,
		Password:     config.RedisConf.Password,
		DB:           config.RedisConf.DB,
		PoolSize:     config.RedisConf.PoolSize,
		DialTimeout:  time.Second * time.Duration(config.RedisConf.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(config.RedisConf.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(config.RedisConf.WriteTimeout),
	}
	log.Debugf("")
	r = redis.NewClient(&opts)
	_, err = r.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("connection redis ping failed, err: %v\n", err)
	}
	return
}
