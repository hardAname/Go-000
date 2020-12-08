package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ErrHandleServer = errors.New("handle http server closed by invoker")
	ErrUpexpected = errors.New("upexpected error")
)

type serveName string
var  (
	nameServer1 serveName = "server1"
	nameServer2 serveName = "server2"
	nameServer3 serveName = "server3"
)

type ServerStr struct{
	name serveName
	server *http.Server
}

func initServer1() *ServerStr{
	str := &ServerStr{
		name:   nameServer1,
		server: &http.Server{
			Addr:              "127.0.0.1:8081",
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(str.name, " handle request")
		if request.Method == http.MethodGet{
			writer.Write([]byte("page1"))
			writer.WriteHeader(http.StatusOK)
		}
		writer.WriteHeader(http.StatusNotFound)
	})
	str.server.Handler = mux
	return str
}

func initServer2() *ServerStr{
	str := &ServerStr{
		name:   nameServer2,
		server: &http.Server{
			Addr:              "127.0.0.1:8082",
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/test2", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(str.name, " handle request")
		if request.Method == http.MethodGet{
			writer.Write([]byte("page2"))
			writer.WriteHeader(http.StatusOK)
		}
		writer.WriteHeader(http.StatusNotFound)
	})
	str.server.Handler = mux
	return str
}

func initServer3() *ServerStr{
	str := &ServerStr{
		name:   nameServer3,
		server: &http.Server{
			Addr:              "127.0.0.1:8083",
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/test3", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(str.name, " handle request")
		if request.Method == http.MethodGet{
			writer.Write([]byte("page3"))
			writer.WriteHeader(http.StatusOK)
		}
		writer.WriteHeader(http.StatusNotFound)
	})
	str.server.Handler = mux
	return str
}

func main(){
	//closeBySignal()
	closeByError()
}

func closeBySignal(){
	var serveArr []*ServerStr
	serveArr = append(serveArr, initServer1())
	serveArr = append(serveArr, initServer2())
	serveArr = append(serveArr, initServer3())

	ctx, cancel := context.WithCancel(context.Background())

	eg,_ := errgroup.WithContext(ctx)
	for _, iStr := range serveArr {
		str := iStr
		eg.Go(func() error {
			defer fmt.Println(str.name," is closed")
			go func() {
				select {
				case <- ctx.Done():
					str.server.Close()
				}
			}()
			err := str.server.ListenAndServe()
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(ErrHandleServer, fmt.Sprintf(string(str.name) + " is closed, inner error:%v", err))
				}
				cancel()
				return errors.Wrap(ErrUpexpected, fmt.Sprintf(string(str.name) + " server inner error:%v", err))
			}
			return nil
		})
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		break
	case s := <-sigCh:
		fmt.Println("receive signal:", s)
		cancel()
	}
	err := eg.Wait()
	if err != nil{
		if errors.Is(err, ErrHandleServer){
			fmt.Printf("errgroup return:%v\n Stack Trace:\n%+v\n", errors.Cause(err), err)
		}else if errors.Is(err, ErrUpexpected){
			fmt.Printf("errgroup closed unexpected:%v\n Stack Trace:\n%+v\n", errors.Cause(err), err)
		}
	}
}

func closeByError(){
	var serveArr []*ServerStr
	serveArr = append(serveArr, initServer1())
	serveArr = append(serveArr, initServer2())
	serveArr = append(serveArr, initServer3())

	ctx, cancel := context.WithCancel(context.Background())

	eg,_ := errgroup.WithContext(ctx)
	for _, iStr := range serveArr {
		str := iStr
		eg.Go(func() error {
			defer fmt.Println(str.name," is closed")
			go func() {
				select {
				case <- ctx.Done():
					str.server.Close()
				}
			}()
			var err error
			/// 人为制造server3的异常推出
			if str.name == nameServer3{
				time.Sleep(time.Second * 15)
				err = errors.New("测试异常退出")
			}else {
				err = str.server.ListenAndServe()
			}
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(ErrHandleServer, fmt.Sprintf(string(str.name) + " is closed, inner error:%v", err))
				}
				defer cancel()
				return errors.Wrap(ErrUpexpected, fmt.Sprintf(string(str.name) + " server inner error:%v", err))
			}
			return nil
		})
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		break
	case s := <-sigCh:
		fmt.Println("receive signal:", s)
		cancel()
	}
	//for _, str := range serveArr{
	//	str.server.Close()
	//}
	//serveArr[0].server.Close()
	err := eg.Wait()
	if err != nil{
		if errors.Is(err, ErrHandleServer){
			fmt.Printf("errgroup return:%v\n Stack Trace:\n%+v\n", errors.Cause(err), err)
		}else if errors.Is(err, ErrUpexpected){
			fmt.Printf("errgroup closed unexpected:%v\n Stack Trace:\n%+v\n", errors.Cause(err), err)
		}
	}
}
