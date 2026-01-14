package errorcode

import "devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/constant"

func init() {
	registerErrorCode(demoErrorMap)
}

// Demo error
const (
	demoPreCoder = constant.ServiceName + "." + demoModelName + "."

	DemoNotExist = demoPreCoder + "DemoNotExist"
)

var demoErrorMap = errorCode{
	DemoNotExist: {
		description: "Demo不存在",
		cause:       "",
		solution:    "请重新选择Demo",
	},
}
