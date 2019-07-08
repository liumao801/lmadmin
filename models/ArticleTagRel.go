package models

import "time"

type ArticleTagRel struct {
	Id 			int
	Article		*Article		`orm:"rel(fk)"`
	ArticleTag 	*ArticleTag		`orm:"rel(fk)"`
	CreatedAt 	time.Time		`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time		`orm:"auto_now_add;type(datetime)"`
}

func (m *ArticleTagRel) TableName() string {
	return ArticleTagRelTBName()
}