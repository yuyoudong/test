package gorm

import (
	demo "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/gorm/demo/impl"
	"github.com/google/wire"
)

var RepositoryProviderSet = wire.NewSet(
	demo.NewRepo,
)
