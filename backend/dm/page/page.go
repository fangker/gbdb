package page

const (
	PAGE_TYPE_PAGE = 1
)
const(
	PAGE_SIZE = 16384
)


type PageData  [PAGE_SIZE]byte



type Page struct {
	fh       FilHead
	data     *PageData
}

func NewPage(data *PageData) *Page {
	page:= &Page{fh:FilHead{},data:data}
	page.fh.parseFilHeader(page.data)
	return  page
}
