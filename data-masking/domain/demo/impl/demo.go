package impl

import (
	domain "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain/demo"
	repo "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/gorm/demo"
)

type useCase struct {
	repoDemo repo.Repo
}

func NewUseCase(repoDemo repo.Repo) domain.UseCase {
	return &useCase{repoDemo: repoDemo}
}
