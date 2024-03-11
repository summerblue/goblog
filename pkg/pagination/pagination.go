package pagination

import (
	"goblog/pkg/config"
	"goblog/pkg/types"
	"math"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Page 单个分页元素
type Page struct {
	// 链接
	URL string
	// 页码
	Number int
}

// ViewData 同视图渲染的数据
type ViewData struct {
	// 是否需要显示分页
	HasPages bool

	// 下一页
	Next    Page
	HasNext bool

	// 上一页
	Prev    Page
	HasPrev bool

	Current Page

	// 数据库的内容总数量
	TotalCount int64
	// 总页数
	TotalPage int
}

// Pagination 分页对象
type Pagination struct {
	BaseURL string
	PerPage int
	Page    int
	Count   int64
	db      *gorm.DB
}

// New 分页对象构建器
// r —— 用来获取分页的 URL 参数，默认是 page，可通过 config/pagination.go 修改
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// baseURL —— 用以分页链接
// PerPage —— 每页条数，传参为小于或者等于 0 时为默认值  10，可通过 config/pagination.go 修改
func New(r *http.Request, db *gorm.DB, baseURL string, PerPage int) *Pagination {

	// 默认每页数量
	if PerPage <= 0 {
		PerPage = config.GetInt("pagination.perpage")
	}

	// 实例对象
	p := &Pagination{
		db:      db,
		PerPage: PerPage,
		Page:    1,
		Count:   -1,
	}

	// 拼接 URL
	if strings.Contains(baseURL, "?") {
		p.BaseURL = baseURL + "&" + config.GetString("pagination.url_query") + "="
	} else {
		p.BaseURL = baseURL + "?" + config.GetString("pagination.url_query") + "="
	}

	// 设置当前页码
	p.SetPage(p.GetPageFromRequest(r))

	return p
}

// Paging 返回渲染分页所需的数据
func (p *Pagination) Paging() ViewData {

	return ViewData{
		HasPages: p.HasPages(),

		Next:    p.NewPage(p.NextPage()),
		HasNext: p.HasNext(),

		Prev:    p.NewPage(p.PrevPage()),
		HasPrev: p.HasPrev(),

		Current:   p.NewPage(p.CurrentPage()),
		TotalPage: p.TotalPage(),

		TotalCount: p.Count,
	}
}

// NewPage 设置当前页
func (p Pagination) NewPage(page int) Page {
	return Page{
		Number: page,
		URL:    p.BaseURL + strconv.Itoa(page),
	}
}

// SetPage 设置当前页
func (p *Pagination) SetPage(page int) {
	if page <= 0 {
		page = 1
	}

	p.Page = page
}

// CurrentPage 返回当前页码
func (p Pagination) CurrentPage() int {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return 0
	}

	if p.Page > totalPage {
		return totalPage
	}

	return p.Page
}

// Results 返回请求数据，请注意 data 参数必须为 GROM 模型的 Slice 对象
func (p Pagination) Results(data interface{}) error {
	var err error
	var offset int
	page := p.CurrentPage()
	if page == 0 {
		return err
	}

	if page > 1 {
		offset = (page - 1) * p.PerPage
	}

	return p.db.Preload(clause.Associations).Limit(p.PerPage).Offset(offset).Find(data).Error
}

// TotalCount 返回的是数据库里的条数
func (p *Pagination) TotalCount() int64 {
	if p.Count == -1 {
		var count int64
		if err := p.db.Count(&count).Error; err != nil {
			return 0
		}
		p.Count = count
	}

	return p.Count
}

// HasPages 总页数大于 1 时会返回 true
func (p *Pagination) HasPages() bool {
	n := p.TotalCount()
	return n > int64(p.PerPage)
}

// HasNext returns true if current page is not the last page
func (p Pagination) HasNext() bool {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return false
	}

	page := p.CurrentPage()
	if page == 0 {
		return false
	}

	return page < totalPage
}

// PrevPage 前一页码，0 意味着这就是第一页
func (p Pagination) PrevPage() int {
	hasPrev := p.HasPrev()

	if !hasPrev {
		return 0
	}

	page := p.CurrentPage()
	if page == 0 {
		return 0
	}

	return page - 1
}

// NextPage 下一页码，0 的话就是最后一页
func (p Pagination) NextPage() int {
	hasNext := p.HasNext()
	if !hasNext {
		return 0
	}

	page := p.CurrentPage()
	if page == 0 {
		return 0
	}

	return page + 1
}

// HasPrev 如果当前页不为第一页，就返回 true
func (p Pagination) HasPrev() bool {
	page := p.CurrentPage()
	if page == 0 {
		return false
	}

	return page > 1
}

// TotalPage 返回总页数
func (p Pagination) TotalPage() int {
	count := p.TotalCount()
	if count == 0 {
		return 0
	}

	nums := int64(math.Ceil(float64(count) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}

// GetPageFromRequest 从 URL 中获取 page 参数
func (p Pagination) GetPageFromRequest(r *http.Request) int {
	page := r.URL.Query().Get(config.GetString("pagination.url_query"))

	if len(page) > 0 {
		pageInt := types.StringToInt(page)
		if pageInt <= 0 {
			return 1
		}
		return pageInt
	}
	return 0
}
