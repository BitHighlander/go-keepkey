package keepkey

import (
	"fmt"
	"sync"
	"testing"
)

func TestMultiplexingKeepkeys(t *testing.T) {
	var wg sync.WaitGroup
	for _, kk := range kks {
		wg.Add(1)
		go func(kk *Keepkey) {
			defer wg.Done()
			err := kk.WipeDevice()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = kk.LoadDevice([]string{"zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "zoo", "wrong"}, "", "label", false, false)
			if err != nil {

				fmt.Println(err)
				return
			}
			_, err = kk.Ping("Ripple", true, false, false)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = kk.EthereumGetAddress([]uint32{0x0}, true)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = kk.Ping("is", true, false, false)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = kk.Ping("> Ellie", true, false, false)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(kk)
	}
	wg.Wait()
}
