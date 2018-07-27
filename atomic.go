package main

import (
	"sync/atomic"
	"fmt"
)

func main()  {
	var a int32 = 3
	b := atomic.AddInt32(&a,3)
	fmt.Printf("a is %d, b is %d\n",a,b)

}
