package models

// OauthWechat 微信用户
type OauthWechat struct {
	ID       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Unionid  string `gorm:"not null" form:"unionid" json:"unionid"`
	Nickname string `gorm:"not null" form:"nickname" json:"nickname"`
}

// TableName 设置表名
func (model *OauthWechat) TableName() string {
	return "oauth_wechat"
}
