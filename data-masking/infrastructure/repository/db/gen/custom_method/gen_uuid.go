package custom_method

import (
	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/util"
	"gorm.io/gorm"
)

type GenIDMethod struct {
	ID string
}

func (m *GenIDMethod) BeforeCreate(_ *gorm.DB) error {
	if m == nil {
		return nil
	}

	if len(m.ID) == 0 {
		m.ID = util.NewUUID()
	}

	return nil
}
