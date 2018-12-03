package home

type MenuWeb struct {
	Id			int
	Title 		string
	Type 		int8
	ModelId 	int
	ParId 		int
	ListTpl 	string
	ArticleTpl 	string
	Url 		string
	Status 		int8
	Sort 		int8
	Img 		string
	SeoTitle 	string
	SeoDesc 	string
	Content 	string
}

type MenuWebQueryParam struct {
	HomeBaseQueryParam
	TitleLike 	string
	Type 		int
	ParId 		int
	Status 		int8
}

func (m *MenuWeb) TableName() string {
	return MenuWebTBName()
}