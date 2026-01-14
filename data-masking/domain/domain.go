package domain

import (
	demo "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain/demo/impl"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	demo.NewUseCase,
)
