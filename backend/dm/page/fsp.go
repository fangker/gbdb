package page

type FSPage struct {
	FH       *FilHeader
	FSH     *FSHeader
	data     *PageData
}

func NewFSPage(data *PageData) *FSPage {
	fsPage:= &FSPage{FH:&FilHeader{data:data},data:data}
	fsPage.FH.ParseFilHeader(fsPage.data)
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	fsPage.FH.SetSpace(0)
	fsPage.FH.SetOffset(0)
	return  fsPage
}

