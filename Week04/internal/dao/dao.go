package dao

import (
	"Go-000/Week04/internal/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDb)

type Dao interface {
	Close()
	GetName(ctx context.Context, id int64) (*model.User, error)
}

type dao struct {
	db *sql.DB
}

func New(db *sql.DB) (d Dao, cf func(), err error){
	return newDao(db)
}

func newDao(db *sql.DB) (d *dao, cf func(), err error){
	// config ..
	d = &dao{
		db:db,
	}
	cf = d.Close
	return
}

func (d *dao) GetName(ctx context.Context, id int64) (user *model.User, err error){
	// do something
	user, err = d.RawGetName(ctx, id)
	return
}

func (d *dao) Close(){
	fmt.Println("dao is closed")
}

