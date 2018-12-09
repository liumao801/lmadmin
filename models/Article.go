package models

import "time"

type Article struct {
	Id				int
	Title 			string
	Subtitle 		string
	Logo 			string
	Keywords 		string
	Desc 			string
	Content 		byte
	MenuId 			int
	Video 			string
	CommentStatus 	int8
	CommentCount 	int
	IsBack 			int8
	Status 			int8
	CreatedAt 		time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt 		int
}

type ArticleQueryParam struct {
	BaseQueryParam
	TitleLike 		string
	SubtitleLike	string
	KeywordsLike 	string
	MenuId 			int
	IsBack 			int
	Status 			int
}

func (m *Article) TableName() string {
	return ArticleTBName()
}