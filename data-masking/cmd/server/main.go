package main

import (
	"flag"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/log"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/settings"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/conf"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db"
	af_go_frame "github.com/jinguoxing/af-go-frame"
	"github.com/jinguoxing/af-go-frame/core/config"
	"github.com/jinguoxing/af-go-frame/core/config/env"
	"github.com/jinguoxing/af-go-frame/core/config/file"
	"github.com/jinguoxing/af-go-frame/core/transport/rest"
)

var (
	Name = "af_data_masking"
	// Version is the version of the compiled software.
	Version = "1.0"

	confPath string
	addr     string
)

type AppRunner struct {
	App *af_go_frame.App
}

func newApp(hs *rest.Server) *af_go_frame.App {

	return af_go_frame.New(
		af_go_frame.Name(Name),
		af_go_frame.Server(hs),
	)
}

func init() {
	flag.StringVar(&confPath, "confPath", "cmd/server/config", "config path, eg: -conf config.yaml")
	flag.StringVar(&addr, "addr", ":8153", "config path, eg: -addr 0.0.0.0:8153")
}

// @title       data-masking
// @version     0.0
// @description AnyFabric data masking
// @BasePath    /api/data-masking/v1
func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			env.NewSource(),
			file.NewSource(confPath),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if addr != "" {
		bc.Server.Http.Addr = addr
	}

	// region 读取配置信息
	var dBConf struct {
		Database db.Database `json:"database"`
	}
	if err := c.Scan(&dBConf); err != nil { // 数据库配置
		panic(err)
	}

	if err := c.Scan(&settings.ConfigInstance); err != nil {
		panic(err)
	}
	if err := c.Scan(&settings.SwagConfig); err != nil {
		panic(err)
	}
	// if settings.SwagConfig.Doc.Host == "" {
	// 	settings.SwagConfig.Doc.Host = "127.0.0.1:8153"
	// }
	settings.CheckConfigPath()
	// 初始化日志
	log.InitProjectLogger()
	defer log.Flush()

	// 初始化验证器
	// err := form_validator.SetupValidator()
	// if err != nil {
	// 	panic(err)
	// }

	appRunner, cleanup, err := InitApp(bc.Server, &dBConf.Database)
	// appRunner, cleanup, err := InitApp(bc.Server)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//start and wait for stop signal
	if err := appRunner.App.Run(); err != nil {
		panic(err)
	}
}
