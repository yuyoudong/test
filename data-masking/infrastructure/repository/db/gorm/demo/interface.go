package demo

import (
	"context"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/model"
)

type Repo interface {
	Get(ctx context.Context, id string) (*model.Demo, error)
}
