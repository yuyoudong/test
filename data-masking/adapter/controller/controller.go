package controller

import (
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/adapter/controller/demo/v1"
	"github.com/google/wire"
)

// HttpProviderSet ProviderSet is server providers.
var HttpProviderSet = wire.NewSet(NewHttpServer)

var ServiceProviderSet = wire.NewSet(
	demo.NewService,
)
