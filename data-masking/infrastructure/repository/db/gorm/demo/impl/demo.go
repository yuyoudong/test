package impl

import (
	"context"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/gorm/demo"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/model"
)

type repo struct {
	data *db.Data
}

func NewRepo(data *db.Data) demo.Repo {
	return &repo{data: data}
}

func (r *repo) Get(ctx context.Context, id string) (*model.Demo, error) {
	_ = r.data.DB.WithContext(ctx)
	panic("impl me")
}
