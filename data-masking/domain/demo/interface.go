package demo

import (
	"context"
)

type UseCase interface {
	Create(ctx context.Context, req *CreateReqParam) (*CreateRespParam, error)
}

/////////////////// Create ///////////////////

type CreateReqParam struct {
	CreateReqPathParam
	CreateReqQueryParam
	CreateReqBodyParam
}

type CreateReqPathParam struct {
	PId string `json:"pid" uri:"pid" binding:"required,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // P ID，uuid
}

type CreateReqQueryParam struct {
	P1 *int    `json:"p1" form:"p1" binding:"omitempty,min=1,max=128" example:"1"` // 查询参数1
	P2 *string `json:"p2" form:"p2" binding:"required,oneof=e1 e2" example:"e1"`   // 查询参数2
}

type CreateReqBodyParam struct {
	Fields    []*FieldInfo `json:"fields"`
	TableName string       `json:"table_name" binding:"required"`
}
type FieldInfo struct {
	Field       string `json:"field" bindging:"required,fl VerifyReq"`
	ChineseName string `json:"chinese_name" bindging:"required"`
	Sensitive   int    `json:"sensitive" bindging:"required"`
	Classified  int    `json:"classified" bindging:"required"`
	FieldType   string `json:"field_type" bindging:"required"`
}

// type CreateReqBodyParam struct {
// 	ID   int     `json:"id" binding:"required"`
// 	Name *string `json:"name" binding:"required,min=1,max=128" example:"demo_name"` // Demo名称，仅支持中英文、数字、下划线及中划线，前后空格自动去除

// }

// func (p *CreateReqParam) ToModel() *model.Demo {
// 	if p == nil {
// 		return nil
// 	}

// 	res := &model.Demo{
// 		Fields: *p.Fields,
// 		TableName: *p.TableName
// 	}
// 	return res
// }

type CreateRespParam struct {
	ID   string `json:"id" binding:"required,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // DemoID
	Name string `json:"name" binding:"required,max=128" example:"demo_name"`                       // Demo名称
}
