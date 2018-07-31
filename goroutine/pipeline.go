package goroutine

import "fmt"

func PipeLine() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)

	// counter
	//go func() {
	//	for x := 0; x < 100; x++ {
	//		naturals <- x
	//	}
	//	close(naturals)
	//}()
	//
	//go func() {
	//	for x := range naturals {
	//		squarer <- x * x
	//	}
	//	// 关闭通道后，接收完最后一个值、再获取到的就是 0 值
	//	close(squarer)
	//}()
	//
	//for x := range squarer {
	//	fmt.Println(x)
	//}
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	// 关闭通道后，接收完最后一个值、再获取到的就是 0 值
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
