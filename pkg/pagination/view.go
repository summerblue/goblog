package pagination

import (
	"goblog/pkg/logger"
	"strconv"
	"strings"

	"github.com/vcraescu/go-paginator"
)

type (
	// Page 分页元素
	Page struct {
		Numer int
		URL   string
	}
	// ViewData ViewData interface
	ViewData interface {
		// 所有可见页码
		Pages() ([]int, error)
		// 下一页
		Next() Page
		// 上一页
		Prev() (int, error)
		// 最后一页
		Last() (int, error)
		// 当前页码
		Current() (int, error)
	}

	// DefaultView 提供给视图的分页数据
	DefaultView struct {
		Paginator paginator.Paginator

		// 用以控制两边页码数量
		Proximity int
		BaseURL   string
	}
)

// NewViewData 分页视图构造器
func NewViewData(p paginator.Paginator, url string) ViewData {
	var baseURL string
	if strings.Contains(url, "?") {
		baseURL = url + "&page="
	} else {
		baseURL = url + "?page="
	}

	return &DefaultView{
		Paginator: p,
		Proximity: 5,
		BaseURL:   baseURL,
	}
}

// Next 返回下一页码，当前页如果是最后页的话，返回 0
func (v *DefaultView) Next() Page {
	num, err := v.Paginator.NextPage()
	logger.LogError(err)
	return Page{
		Numer: num,
		URL:   v.BaseURL + strconv.Itoa(num),
	}
}

// Prev 返回上一页码，当前页如果是最后页的话，返回 0
func (v *DefaultView) Prev() (int, error) {
	return v.Paginator.PrevPage()
}

// Last 返回最后页码
func (v *DefaultView) Last() (int, error) {
	return v.Paginator.PageNums()
}

// Current 返回当前页码
func (v *DefaultView) Current() (int, error) {
	return v.Paginator.Page()
}

// Pages 所有可见页码
func (v *DefaultView) Pages() ([]int, error) {
	var items []int
	hasPages, err := v.Paginator.HasPages()
	if err != nil {
		return nil, err
	}

	if !hasPages {
		return items, nil
	}

	items = make([]int, 0)
	length := v.Proximity * 2
	pn, err := v.Paginator.PageNums()
	if err != nil {
		return nil, err
	}

	if pn < length {
		pn, err := v.Paginator.PageNums()
		if err != nil {
			return nil, err
		}

		length = pn
	}

	proximityLeft := length / 2
	proximityRight := (length / 2) - 1
	if length%2 != 0 {
		proximityRight = proximityLeft
	}

	page, err := v.Paginator.Page()
	if err != nil {
		return nil, err
	}

	start := page - proximityLeft
	end := page + proximityRight
	if start <= 0 {
		start = 1
		end = length
	}

	for page = start; page <= end; page++ {
		items = append(items, page)
	}

	return items, nil
}
