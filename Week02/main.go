package main

import (
	"Go-000/Week02/dao"
	"fmt"
	"github.com/pkg/errors"
)

func invokeDao(callNum int) {
	d := dao.DaoA{}
	target, err := d.FindTargetByNum(callNum)
	if err != nil{
		fmt.Printf("DaoA num:%d error:%T %v\n",callNum, errors.Cause(err), errors.Cause(err))
		fmt.Printf("Stack Trace:\n%+v\n", err)
		return
	}
	fmt.Printf("DaoA num:%d, got result:%s\n",callNum, target.Content)
}

func main(){
	invokeDao(2)
	fmt.Println("-----------")
	invokeDao(1)
}
