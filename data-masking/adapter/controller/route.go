package controller

import (
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/adapter/controller/demo/v1"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(r *gin.Engine) error
}

type Router struct {
	DemoApi *demo.Service
}

func (r *Router) Register(engine *gin.Engine) error {
	r.RegisterApi(engine)
	return nil
}

func (r *Router) RegisterApi(engine *gin.Engine) {
	dataMaskingRouter := engine.Group("/api/data-security/v1")
	{

		{
			demoRouter := dataMaskingRouter.Group("/data-masking")

			demoRouter.POST("/sql-masking", r.DemoApi.Create)
		}
	}
}
