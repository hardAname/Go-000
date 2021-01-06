package Week06

import (
	"sync/atomic"
	"time"
)

type Bucket struct{
	StartTime int64
	num  int64
}

func (b *Bucket) Inc(i int64){
	atomic.AddInt64(&b.num, i)
}

func (b *Bucket) Get() int64{
	return b.num
}

func (b *Bucket) Reset(){
	_ = atomic.SwapInt64(&b.StartTime, 0)
	_ = atomic.SwapInt64(&b.num, 0)
}

type Circular struct{
	numBucket int64
	intervalTime int64
	totalTime int64
	bucketArr	[]*Bucket
	blockCh		chan struct{}
}

func NewCircular(size int64, interval int64, currentTime int64) *Circular{
	var bArr []*Bucket
	for i := int64(0); i < size; i++ {
		bArr = append(bArr, &Bucket{
			StartTime: currentTime + interval * i,
		})
	}
	return &Circular{
		numBucket:    size,
		intervalTime: interval,
		totalTime: size * interval,
		bucketArr:    bArr,
		blockCh: 	make(chan struct{}, 1),
	}
}

func (cl *Circular) GetCurrentBucket() (bucket *Bucket){
	currentTime := time.Now().Unix()/1000000
	select {
	case cl.blockCh <- struct{}{}:
		for i := int64(0); i < cl.numBucket; i++{
			bucket = cl.bucketArr[i]
			if currentTime < (bucket.StartTime + cl.intervalTime){
				return
			}else if (currentTime - (bucket.StartTime + cl.intervalTime)) > cl.totalTime{
				cl.reset()
				return cl.GetCurrentBucket()
			}else{
				return bucket
			}
		}
	default:
		bucket = cl.bucketArr[cl.numBucket - 1]
		if bucket.StartTime != 0{
			return bucket
		}else{
			time.Sleep(time.Second)
			return cl.GetCurrentBucket()
		}
	}
	return
}

func (cl *Circular) reset(){
	currentTime := int64(time.Now().UnixNano()/1000000)
	for i := int64(0); i < cl.numBucket; i++{
		cl.bucketArr[i].Reset()
		cl.bucketArr[i].StartTime = currentTime + i * cl.intervalTime
	}
}

func (cl *Circular) GetSum() (sum int64){
	for i := int64(0); i < cl.numBucket; i++{
		sum += cl.bucketArr[i].num
	}
	return
}
