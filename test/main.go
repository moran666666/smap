package main

import (
	"fmt"
	"sync"

  "github.com/moran666666/smap"
)

func main() {
	var waitGroup sync.WaitGroup

	safeMap := smap.NewSmap()

	// safeMap.Set(k, v)
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			safeMap.Set(i, i)
		}(i)
	}
	waitGroup.Wait()

	// safeMap.Delete(k)
	safeMap.Delete(6)

	// elt := safeMap.Get(k)
	for j := 0; j < 10; j++ {
		waitGroup.Add(1)
		go func(j int) {
			defer waitGroup.Done()
			elt := safeMap.Get(j)
			if elt == nil {
				fmt.Println("sMap[", j, "]:", "nil")
			} else {
				fmt.Println("sMap[", j, "]:", elt)
			}
		}(j)
	}
	waitGroup.Wait()

	// safeMap.Iter(chan interface{})
	iterCh := make(chan interface{})
	safeMap.Iter(iterCh)
	for iter := range iterCh {
		kv := iter.(*smap.KeyValuePair)
		k := kv.Key.(int)
		v := kv.Value.(int)
		fmt.Println("Key:", k, "Value:", v)
	}

	// safeMap.Len()
	length := safeMap.Len()
	fmt.Println("sMap Len():", length)
}
