package main

//func main() {
//	// 当下全局对象
//	cfg := &Config{}
//
//	go func() {
//		i := 0
//		for {
//			i++
//			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
//		}
//	}()
//
//	var wg sync.WaitGroup
//	for n := 0; n < 4; n++ {
//		wg.Add(1)
//		go func() {
//			for i := 0; i < 100; i++ {
//				// 外部在读的时候,可能内部又进行了写,因为slice读并不是原子的
//				// 可能在你读的同时,那边又写入了新的值
//				fmt.Printf("%v\n", cfg)
//			}
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}
