/**
 * @Time : 2019/7/2 19:19 
 * @Author : liumao801 
 * @File : paginator
 * @Software: GoLand
 * Readme : https://beego.me/docs/mvc/view/page.md
 *
 * 分页处理结构体
 *
 */
package utils

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Paginator struct {
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

func (p *Paginator) Nums() int64 {
	return p.nums
}

func (p *Paginator) SetNums(nums interface{}) {
	p.nums, _ = ToInt64(nums)
}

func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i, _ := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i, _ := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i, _ := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}

func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	values := link.Query()
	if page == 1 {
		values.Del("p")
	} else {
		values.Set("p", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}

func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return
}

func (p *Paginator) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

func (p *Paginator) PageLinkLast() (link string) {
	return p.PageLink(p.PageNums())
}

func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

func NewPaginator(req *http.Request, per int, nums interface{}) *Paginator {
	p := Paginator{}
	p.Request = req
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetNums(nums)
	return &p
}


/**
// 分页调用方法 以 ArticleController 为例
// @params totals 总数据条数
func (c *ArticleController) SetPaginator(totals int64) *utils.Paginator {
	per := 3 // 每页显示几条数据
	p := utils.NewPaginator(c.Ctx.Request, per, totals)
	c.Data["paginator"] = p
	return p
}

// 设置分页参数
func (c *ArticleController) setPageOffset(params *models.ArticleQueryParam) {
	params.Limit = limit
	page, _ := c.GetInt("p", 1)
	offset := limit * (page - 1)
	params.Offset = int64(offset)
}



// model 中查询当前页
query.Limit(params.Limit, params.Offset).All(&data)




// 前端分页代码
	<!-- 前端分页代码 start -->
{{if gt .paginator.PageNums 1}}
	<div class="col-xs-12 col-sm-12 col-md-12 col-lg-12 padding-v10px line-right">
		<ul class="pagination pagination-sm" style="margin: 0;">
		{{if .paginator.HasPrev}}
			<li><a href="{{.paginator.PageLinkFirst}}" title="首页">&lt;&lt;</a></li>
			<li><a href="{{.paginator.PageLinkPrev}}" title="上一页">&lt;</a></li>
		{{else}}
			<li class="disabled"><a>&lt;&lt;</a></li>
			<li class="disabled"><a>&lt;</a></li>
		{{end}}
		{{range $index, $page := .paginator.Pages}}
			<li{{if $.paginator.IsActive .}} class="active"{{end}}>
				<a href="{{$.paginator.PageLink $page}}" title="第{{$page}}页">{{$page}}</a>
			</li>
		{{end}}
		{{if .paginator.HasNext}}
			<li><a href="{{.paginator.PageLinkNext}}" title="下一页">&gt;</a></li>
			<li><a href="{{.paginator.PageLinkLast}}" title="最后页">&gt;&gt;</a></li>
		{{else}}
			<li class="disabled"><a>&gt;</a></li>
			<li class="disabled"><a>&gt;&gt;</a></li>
		{{end}}
		</ul>
	</div>
{{end}}
	<!-- 前端分页代码 end -->

*/