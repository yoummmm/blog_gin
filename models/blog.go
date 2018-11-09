package models

import "blog_gin/database"

type Blog struct {
	id          int    `json:"id" form:"id"`
	title       string `json:"title" form:"title"`
	abstr       string `json:"abstr" form:"abstr"`
	category_id string `json:"category_id" form:"category_id"`
	key         string `json:"key" form:"key"`
	sorted      string `json:"sorted" form:"sorted"`
	author      string `json:"author" form:"author"`
	content_id  string `json:"content_id" form:"content_id"`
}

func (B *Blog) ListBlogs(pagenum int) (res []Blog,err error) {
	  res = make([]Blog,0)
		database.Engine.Distinct("id","title","category_id","author").Limit(10,pagenum).Find(&res)
		return
}
