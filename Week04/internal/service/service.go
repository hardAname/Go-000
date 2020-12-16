package service

import (
	"Go-000/Week04/api"
	"Go-000/Week04/internal/dao"
	"context"
	"fmt"
)

type Service struct {
	dao dao.Dao
}

func New(d dao.Dao) (s *Service, cf func(), err error){
	s = &Service{dao:d}
	cf = s.Close
	return
}

func (svr *Service) SayHello(ctx context.Context,req *api.HelloReq) (resp *api.HelloResp, err error){
	u, err := svr.dao.GetName(ctx, req.Id)
	if err != nil{
		return nil, err
	}
	resp = &api.HelloResp{Content:u.Name}
	return
}

func (svr *Service) Close(){
	fmt.Println("Service is Closed")
}

