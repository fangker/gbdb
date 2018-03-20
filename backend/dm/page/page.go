package page

const (
	PAGE_TYPE_FSP   = 0
	PAGE_TYPE_INODE = 1
	PAGE_TYPE_PAGE  = 2
	PAGE_TYPE_INDEX = 3
)
const (
	PAGE_SIZE = 16384
)

type PageData [PAGE_SIZE]byte

type Page struct {
	FH   FilHeader
	data *PageData
}

func NewPage(data *PageData) *Page {
	page := &Page{FH: FilHeader{}, data: data}
	page.FH.ParseFilHeader(page.data)
	return page
}
