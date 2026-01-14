package repository

import (
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/conf"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db"
)

type IData interface {
	NewData(c *conf.Data) (*db.Data, func(), error)
}
