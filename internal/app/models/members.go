package models

type MemberModel struct {
	BaseModel
	Userid string `json:"userid" gorm:"primaryKey;column:userid;type:varchar(64);index:userid_index;comment:userid"`
}

// TableName 企业微信员工
func (MemberModel) TableName() string {
	return "members"
}
