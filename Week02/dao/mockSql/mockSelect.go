package mockSql
/// 模拟通用数据库操作
import (
	"errors"
)

var ErrNoRows  = errors.New("ErrNoRows")

type Db struct{
	addr string
}

type Rows interface {
	Next() bool
	Scan() string
}

type rows struct{
	result string
}

func (r *rows) Next() bool{
	if len(r.result) > 0{
		return true
	}
	return false
}

func (r *rows) Scan() string{
	return r.result
}

func NewDb(addr string) *Db{
	return &Db{addr:addr}
}

func (db *Db) Query(query string, args... interface{}) (Rows, error){
	if query != "select * from table1" || len(args) > 0{
		return nil, ErrNoRows
	}
	return &rows{result:"you got it"}, nil
}
