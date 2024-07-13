package db

type LinkModel struct {
	Id     string
	Name   string
	Target string
	Ctime  string
	Mtime  string
	Visits int
}

// All possible db operations should be in a dal
type LinkDAL interface {
	GetById(id string) (LinkModel, error)
	GetAll() ([]LinkModel, error)
	Insert(link LinkModel) (LinkModel, error)
}

type linkDalImpl struct {
	db []LinkModel
}

func NewLinkDAL() (LinkDAL, error) {
	dal := &linkDalImpl{
		db: make([]LinkModel, 0),
	}

	return dal, nil
}

func (ld *linkDalImpl) GetAll() ([]LinkModel, error) {
	return ld.db, nil
}

func (ld *linkDalImpl) GetById(id string) (LinkModel, error)     { panic("TODO") }
func (ld *linkDalImpl) Insert(link LinkModel) (LinkModel, error) { panic("TODO") }
