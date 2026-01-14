package impl

import (
	"context"

	domain "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain/demo"
)

func (u *useCase) Create(ctx context.Context, req *domain.CreateReqParam) (*domain.CreateRespParam, error) {
	_, _ = u.repoDemo.Get(ctx, "demo id")
	panic("impl me")
}
