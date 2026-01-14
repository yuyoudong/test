//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/adapter/controller"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/domain"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/conf"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db"
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/infrastructure/repository/db/gorm"
	"github.com/google/wire"
)

var appRunnerSet = wire.NewSet(wire.Struct(new(AppRunner), "*"))

func InitApp(*conf.Server, *db.Database) (*AppRunner, func(), error) {
	// func InitApp(*conf.Server) (*AppRunner, func(), error) {
	panic(wire.Build(
		controller.HttpProviderSet,
		controller.RouterSet,
		controller.ServiceProviderSet,
		domain.ProviderSet,
		gorm.RepositoryProviderSet,
		// db.NewData,
		newApp,
		appRunnerSet))
}
