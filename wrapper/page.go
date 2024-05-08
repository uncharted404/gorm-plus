package wrapper

type Page struct {
	Page     int
	PageSize int
}

func NewPage(page, pageSize int) *Page {
	return &Page{Page: page, PageSize: pageSize}
}
