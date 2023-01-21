package main

import (
	"fmt"
	"sync"

	clientone "gitlab.com/chicho69-cesar/singleton/client-one"
	clienttwo "gitlab.com/chicho69-cesar/singleton/client-two"
	"gitlab.com/chicho69-cesar/singleton/singleton"
)

func main() {
	wg := sync.WaitGroup {}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			clientone.IncrementAge()
		}()
		
		go func()  {
			defer wg.Done()
			clienttwo.IncrementAge()
		}()
	}

	wg.Wait()

	p := singleton.GetInstance()
	age := p.GetAge()

	fmt.Println(age)
}
