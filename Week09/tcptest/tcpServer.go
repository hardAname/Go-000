package tcptest

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"sync"
)

type TcpServer struct{
	listener net.Listener
	ctx context.Context
	cancel func()
	once sync.Once
}

type Client struct{
	writeChan chan *Msg
	conn net.Conn
	once sync.Once
	ctx  context.Context
	cancel func()
}

type Msg struct{
	Cmd int32
	Data []byte
}

func NewTcpServer(addr string) (ts *TcpServer, err error){
	lis, err := net.Listen("tcp", addr)
	if err != nil{
		err = errors.Wrap(err, "net.Listen error")
		return
	}
	ts = &TcpServer{
		listener: lis,
		once:   sync.Once{},
	}
	ts.ctx, ts.cancel = context.WithCancel(context.Background())
	go func() {
		for{
			conn, err := lis.Accept()
			if err != nil{
				fmt.Println("lis.Accept error:", err)
				return
			}
			go ts.handleConn(conn)
		}
	}()
	return
}

func (ts *TcpServer) ShutDown(){
	ts.once.Do(func() {
		ts.listener.Close()
		ts.cancel()
	})
}

func (ts *TcpServer) handleConn(conn net.Conn){
	cl := &Client{
		writeChan: make(chan *Msg, 10),
		conn:      conn,
		once:      sync.Once{},
	}
	ctx, cancel := context.WithCancel(ts.ctx)
	defer func() {
		fmt.Println("handleConn return ...")
		cancel()
		cl.Close()
	}()
	go cl.handleWrite(ctx)
	var data []byte
	scan := bufio.NewScanner(conn)
	for scan.Scan(){
		data = scan.Bytes()
		fmt.Println("receive data:", string(data))
		cl.writeChan <-&Msg{
			Cmd:  1,
			Data: []byte("message from server\n"),
		}
	}
	fmt.Println("scan error:", scan.Err())
}

func (cl *Client) handleWrite(ctx context.Context){
	defer cl.Close()
	for {
		select {
		case msg := <-cl.writeChan:
			if msg == nil {
				fmt.Println("handleWrite return chan...")
				return
			}
			fmt.Println("to write msg...")
			cl.conn.Write(msg.Data)
		case <-ctx.Done():
			fmt.Println("handleWrite return ctx...")
			return
		}
	}
}

func (cl *Client) Close(){
	cl.once.Do(func() {
		cl.conn.Close()
	})
}
