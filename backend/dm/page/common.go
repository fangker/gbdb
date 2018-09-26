package page

type Pos struct {
	page   uint32
	offset uint16
}

func NPos(p uint32, o uint16) Pos {
	return Pos{page: p, offset: o}
}

func (p *Pos) ISNil() bool {
 if p.page==0&&p.offset==0 {
 	return true
 }
 return false
}
