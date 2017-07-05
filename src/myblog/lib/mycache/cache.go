package mycache

import (
	"time"
    "fmt"
)

type cacher interface {
	Refresh() error
	Rinterval() int64
}

type cache struct {
	cacher
	RefreshCount  int64
	LastRefresh   int64
	LastFailed    int64
	LastError     error
}

//var cache_c = make(chan cacher)
var cache_c chan cacher

func (self * cache)Load() {
	now := time.Now().Unix()
	if now - self.LastRefresh < self.Rinterval() {
		return
	}

	self.RefreshCount = self.RefreshCount + 1
	self.LastRefresh = now

	err := self.Refresh()
	if err != nil {
		self.LastFailed = now
		self.LastError = err
	}
}

func AddCache(x cacher){
	cache_c <- x
}

func StopCache(){
	close(cache_c)
}

func cacheLoop(){
	var cachelist []*cache
	timer := time.NewTicker(5 * time.Second)
	//check := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <- timer.C:
			for i,_ := range cachelist {
				go cachelist[i].Load()
			}
		case x := <-cache_c:
			y :=  &cache{
				cacher: x,
				RefreshCount: 0,
				LastRefresh: time.Now().Unix(),
			}
            cachelist = append(cachelist,y)
		}
	}
}

func init(){
    cache_c = make(chan cacher)
	go cacheLoop()
}



