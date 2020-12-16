package dao

import (
	"Go-000/Week04/internal/model"
	"database/sql"
	"fmt"
	"context"
)

func NewDb() (db *sql.DB, cf func(), err error){
	// db, err = xxx
	cf = func() {
		fmt.Println("db is closed")
	}
	return
}

func (d *dao) RawGetName(ctx context.Context, id int64) (user *model.User, err error){
	user = &model.User{
		Id:   id,
		Name: "mock const name",
	}
	return
}


