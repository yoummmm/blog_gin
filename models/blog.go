package models

// Blog 博客
type Blog struct {
	// 主键
	ID int16 `json:"id" form:"id" xorm:"pk id"`
	// 标题
	Title string `json:"title" form:"title"`
	//描述
	Abstr string `json:"abstr" form:"abstr"`
	//类型
	Category_id string `json:"category_id" form:"category_id"`
	
	Author string `json:"author" form:"author"`
}

// TableName set table
func (Blog) TableName() string {
	return "blog"
}
