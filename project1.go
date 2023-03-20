package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mt sync.Mutex

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go random(i, data1, &wg)
		go random(i, data2, &wg)
	}

	wg.Wait()

	fmt.Println("===============")

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go random2(i, data1, &wg, &mt)
		wg.Add(1)
		go random2(i, data2, &wg, &mt)
	}

	wg.Wait()
}

func random(index int, data interface{}, wg *sync.WaitGroup) {
	fmt.Printf("%v %d\n", data, index)
	wg.Done()
}

func random2(index int, data interface{}, wg *sync.WaitGroup, mt *sync.Mutex) {
	mt.Lock()
	fmt.Printf("%v %d\n", data, index)
	mt.Unlock()
	wg.Done()
}
