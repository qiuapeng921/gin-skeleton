package request

type EnterpriseBaseReq struct {
	EnterpriseId int64 `json:"enterprise_id" binding:"required" comment:"企业ID"`
}

type PageReq struct {
	Page int `json:"page" binding:"min=1" comment:"页码"`
	Size int `json:"size" binding:"min=10,max=1000" comment:"分页数"`
}
