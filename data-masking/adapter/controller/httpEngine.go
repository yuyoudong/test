package controller

import (
	// "time"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/cmd/server/docs"
	// "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/log"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/settings"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/conf"
	"github.com/gin-gonic/gin"

	// "github.com/jinguoxing/af-go-frame/core/middleware/ginMiddleWare"
	"github.com/jinguoxing/af-go-frame/core/transport/rest"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/zeromicro/go-zero/core/logx"
)

func NewHttpServer(c *conf.Server, r IRouter) *rest.Server {

	engine := NewHttpEngine(r)

	httpSrv := rest.NewServer(engine, rest.Address(c.Http.Addr))

	return httpSrv
}

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine(r IRouter) *gin.Engine {
	// 设置为Release，为的是默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	// 默认启动一个Web引擎
	app := gin.New()

	// 默认注册recovery中间件
	app.Use(gin.Recovery())

	// writer := log.NewZapWriter()
	// logx.SetWriter(writer)

	// app.Use(ginMiddleWare.GinZap(writer, time.RFC3339, false))
	// app.Use(ginMiddleWare.RecoveryWithZap(writer, true))

	docs.SwaggerInfo.Host = settings.SwagConfig.Doc.Host
	docs.SwaggerInfo.Version = settings.SwagConfig.Doc.Version

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 业务绑定路由操作
	err := r.Register(app)
	if err != nil {
		panic(err)
	}

	// 返回绑定路由后的Web引擎
	return app
}
