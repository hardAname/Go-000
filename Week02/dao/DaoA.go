package dao

import (
	"Go-000/Week02/dao/mockSql"
	"github.com/pkg/errors"
	"strconv"
)
// 模拟dao层某个业务模块调用的结构
type DaoA struct{
	db 	*mockSql.Db
}

type TargetStr struct {
	Content string
}

func (dao *DaoA) Init(addr string){
	dao.db = mockSql.NewDb(addr)
}
// 特定的sql行为
func (dao *DaoA) FindTargetByNum(num int) (*TargetStr, error){
	var rows mockSql.Rows
	var err error
	switch num {
	case 1:
		rows, err = dao.db.Query("select * from table1")
	default:
		rows, err = dao.db.Query("执行某些特定的sql")
	}
	if err != nil {
		return nil, errors.Wrap(err, "FindTargetByNum:" + strconv.Itoa(num))
	}
	return &TargetStr{Content:rows.Scan()}, nil
}
