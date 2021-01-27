package tcptest

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"time"
)

func RunClient(addr string) (closeFunc func(),err error){
	conn, err := net.Dial("tcp", addr)
	if err != nil{
		 err = errors.Wrap(err, "net.Dial error")
		 return
	}
	subctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		scan := bufio.NewScanner(conn)
		var data []byte
		for scan.Scan(){
			data = scan.Bytes()
			fmt.Println("receive data:", string(data))
		}
		fmt.Println("scan error:", scan.Err())
	}()
	go func() {
		timer := time.NewTimer(time.Second * 3)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("send data ...")
				_, err = conn.Write([]byte("message from client\n"))
				if err != nil{
					fmt.Println("write error...")
					return
				}
				timer.Reset(time.Second * 3)
			case <-subctx.Done():
				fmt.Println("cancel ...")
				conn.Close()
				return
			}
		}
	}()
	return func() {
		cancel()
	}, nil
}
