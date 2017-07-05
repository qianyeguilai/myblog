package mycache

import (
    "time"
    "testing"
)

type Tcache struct {
    k int
}

func (c * Tcache)Refresh() error{
    c.k = 1
    return nil
}

func (c * Tcache)Rinterval()int64 {
    return 1
}

func Test_Cache(t * testing.T){
    var x Tcache
    AddCache(&x)
    time.Sleep(10 * time.Second)
    if x.k != 1 {
        t.FailNow()
    }
}
