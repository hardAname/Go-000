package Week06_test

import (
	"Go-000/Week06"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewCircular(t *testing.T) {
	circular := Week06.NewCircular(10, 100, time.Now().UnixNano()/1000000 )
	loops := 6
	wg := sync.WaitGroup{}
	wg.Add(loops)
	closeCh := make(chan struct{})
	for i := 0; i < loops; i++{
		go func() {
			defer wg.Done()
			for j := 0; j < 11; j++{
				bucket := circular.GetCurrentBucket()
				for i := 0; i < 100000; i++{
					bucket.Inc(1)
				}
				time.Sleep(time.Second)
			}
		}()
	}

	go func() {
		for {
			select {
			case <-closeCh:
				return
			default:
				time.Sleep(time.Second * 1)
				fmt.Println("sum:",circular.GetSum())
			}
		}

	}()
	wg.Wait()
	close(closeCh)
}
