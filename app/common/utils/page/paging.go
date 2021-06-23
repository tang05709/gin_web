package page

import (
	"festival/app/request"
	"math"
)

type Paging struct {
	PageNum   int `json:"page"`      //当前页
	Pagesize  int `json:"pagesize"`  //每页条数
	Total     int `json:"total"`     //总条数
	PageCount int `json:"pageCount"` //总页数
	StartNum  int `json:"-"`         //起始行
}

//创建分页
func CreatePaging(pageNum, pagesize, total int) *Paging {
	if pageNum < 1 {
		pageNum = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	page_count := math.Ceil(float64(total) / float64(pagesize))
	strat_num := pagesize * (pageNum - 1)
	paging := new(Paging)
	paging.PageNum = pageNum
	paging.Pagesize = pagesize
	paging.Total = total
	paging.PageCount = int(page_count)
	paging.StartNum = strat_num
	return paging
}

func ToPaging(req *request.PageReq) *Paging {
	if req == nil {
		req = &request.PageReq{Page: 1, PageSize: 10}
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	// page_count := math.Ceil(float64(total) / float64(pagesize))
	strat_num := req.PageSize * (req.Page - 1)
	paging := new(Paging)
	paging.PageNum = req.Page
	paging.Pagesize = req.PageSize
	paging.Total = 0
	paging.PageCount = 0
	// paging.Total = total
	// paging.PageCount = int(page_count)
	paging.StartNum = strat_num
	return paging
}

func (pag Paging) CalcPage(total int64) *Paging {
	page_count := math.Ceil(float64(total) / float64(pag.Pagesize))
	pag.Total = int(total)
	pag.PageCount = int(page_count)
	return &pag
}
