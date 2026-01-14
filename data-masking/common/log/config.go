package log

import (
	"fmt"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/settings"
	"github.com/jinguoxing/af-go-frame/core/logx/zapx"
)

const appLoggerName = "data_masking"
const requestLoggerName = "data_masking_request"

var logger zapx.Logger
var logPath string

func InitProjectLogger() {
	logPath = settings.ConfigInstance.Config.LogPath
	logger = newProjectLogger(appLoggerName)
}

func newProjectLogger(loggerName string) zapx.Logger {
	infoConfig := zapx.CoreConfig{
		RotateSize:  10 * zapx.MB,
		Destination: fmt.Sprintf("%s/info.log", logPath),
		Format:      zapx.ConsoleFormat,
		LogLevel:    zapx.DefaultInfoLevel(),
	}
	errorConfig := zapx.CoreConfig{
		RotateSize:  10 * zapx.MB,
		Destination: fmt.Sprintf("%s/error.log", logPath),
		Format:      zapx.JsonFormat,
		LogLevel:    zapx.DefaultErrorLevel(),
	}
	consoleConfig := zapx.CoreConfig{
		RotateSize:  10 * zapx.MB,
		Destination: zapx.ConsoleDestination,
		Format:      zapx.ConsoleFormat,
		LogLevel:    zapx.DefaultConsoleLevel(),
	}
	options := zapx.Options{
		Name:            loggerName,
		EnableCaller:    true,
		CallerSkip:      zapx.ProjectCallerSkip,
		Development:     true,
		StacktraceLevel: zapx.DisableLevel,
		CoreConfigs:     []zapx.CoreConfig{infoConfig, errorConfig, consoleConfig},
	}
	return options.NewZapLogger()
}

func NewZapWriter() *zapx.ZapWriter {
	infoConfig := zapx.CoreConfig{
		RotateSize:  10 * zapx.MB,
		Destination: fmt.Sprintf("%s/requerst.log", logPath),
		Format:      zapx.ConsoleFormat,
		LogLevel:    zapx.DefaultConsoleLevel(),
	}
	consoleConfig := zapx.CoreConfig{
		Destination: zapx.ConsoleDestination,
		Format:      zapx.ConsoleFormat,
		LogLevel:    zapx.DefaultConsoleLevel(),
	}
	options := zapx.Options{
		Name:            requestLoggerName,
		EnableCaller:    true,
		CallerSkip:      zapx.ProjectCallerSkip + 1,
		Development:     true,
		StacktraceLevel: zapx.DisableLevel,
		CoreConfigs:     []zapx.CoreConfig{infoConfig, consoleConfig},
	}
	return zapx.NewZapWriter(*options.NewZapLogger())
}
