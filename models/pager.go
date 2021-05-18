package models

//分页对象
type Pager struct {
	//当前页码
	Page int
	//每页大小
	Pagesize int
	// 文章总数量
	Totalnum int
	//url
	urlpath string
	//总页数，通过每页大小和文章总数计算出来的
}

//分页的构造函数
func NewPager(page, pagesize, totalnum int, urlpath string)*Pager{
	pager := new(Pager)
	pager.Page = page
	pager.Pagesize = pagesize
	pager.Totalnum = totalnum
	pager.urlpath = urlpath
	return pager
}

//各种set方法
func (pager * Pager) SetPage(page int) {
	pager.Page = page
}

func (pager * Pager) SetPagesize(pagesize int) {
	pager.Pagesize = pagesize
}

func (pager * Pager) SetTotalnum(totalnume int) {
	pager.Totalnum = totalnume
}

func (pager * Pager) SetUrlpath(urlpath string) {
	pager.urlpath = urlpath
}


