package tcptest_test

import (
	"Go-000/Week09/tcptest"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNewTcpServer(t *testing.T) {
	ts, err := tcptest.NewTcpServer("127.0.0.1:12306")
	if err != nil{
		t.Fatalf("NewTcpServer error: %v\n", err)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	fmt.Println("receive signal:", s)
	ts.ShutDown()
	time.Sleep(time.Second)
}
