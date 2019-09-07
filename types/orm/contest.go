package orm

import (
	"time"

	"github.com/go-xorm/xorm"
)

// Contest example
type Contest struct {
	ID int `xorm:"not null pk autoincr 'id'"`

	BeginTime time.Time `xorm:"not null 'begin_time'"`
	EndTime   time.Time `xorm:"not null 'end_time'"`

	/* todo: redis */
}

// TableName return the table name
func (obj *Contest) TableName() string {
	return "contest"
}

// GetSliceWithPredict return the slice of object with reserving the space of n Contest
func (obj *Contest) GetSliceWithPredict(n int) interface{} {
	return make([]Contest, 0, n)
}

// GetSlice return the slice of object
func (obj *Contest) GetSlice() interface{} {
	return make([]Contest, 0)
}

// Insert into Engine
func (obj *Contest) Insert() (int64, error) {
	return x.Insert(obj)
}

// Delete from Engine
func (obj *Contest) Delete() (int64, error) {
	return x.ID(obj.ID).Delete(obj)
}

// Update to Engine
func (obj *Contest) Update() (int64, error) {
	return x.ID(obj.ID).Update(obj)
}

// Query from Engine
func (obj *Contest) Query() (bool, error) {
	return x.Get(obj)
}

// Contester Extend the Engine operation
type Contester struct {
}

func NewObjector() (*Contester, error) {
	return new(Contester), nil
}

// ContestSession Extend the Engine operation
type ContestSession xorm.Session

// Query return the code with Property property
func (objx *Contester) Query(property int) (*Contest, error) {
	obj := new(Contest)
	obj.ID = property
	has, err := x.Get(obj)
	if has {
		return obj, nil
	}
	return nil, err
}

// Inserts many objects
func (objx *Contester) Inserts(objs []Contest) (int64, error) {
	return x.Insert(objs)
}

// Querys with conditions
func (objx *Contester) Querys(objs []Contest, conds ...interface{}) error {
	return x.Find(&objs, conds...)
}

// ColsQuerys with conditions with specifying columns
func (objx *Contester) ColsQuerys(objs []Contest, cols ...string) error {
	return x.Cols(cols...).Find(&objs)
}

// Where provIDes custom query condition.
func (objx *Contester) Where(query interface{}, args ...interface{}) *ContestSession {
	return (*ContestSession)(x.Where(query, args...))
}

// Where provIDes custom query condition.
func (objx *ContestSession) Where(query interface{}, args ...interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).Where(query, args...))
}

// And provIDes custom query condition.
func (objx *ContestSession) And(query interface{}, args ...interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).And(query, args...))
}

// Or provIDes custom query condition.
func (objx *ContestSession) Or(query interface{}, args ...interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).Or(query, args...))
}

// ID provIDes custom query condition.
func (objx *ContestSession) ID(query interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).ID(query))
}

// NotIn provIDes custom query condition.
func (objx *ContestSession) NotIn(query string, args ...interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).NotIn(query, args...))
}

// In provIDes custom query condition.
func (objx *ContestSession) In(query string, args ...interface{}) *ContestSession {
	return (*ContestSession)(((*xorm.Session)(objx)).In(query, args...))
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (objx *ContestSession) Find(conds ...interface{}) ([]Contest, error) {
	objs := make([]Contest, 0)
	err := ((*xorm.Session)(objx)).Find(&objs, conds...)
	return objs, err
}
