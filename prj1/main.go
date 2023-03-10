package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	Working()
}

func Working() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go DOSomething(i, c)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func DOSomething(times int, c chan int) {
	time.Sleep(1 * time.Second)
	c <- times
}




//---used for the testing purpose----

/*func main() {
	fmt.Println("Hello, 世界")
	Working()
}

func Working() {
	for i := 0; i < 5; i++ {
		DOSomething(i)
	}
}

func DOSomething(times int) {
	time.Sleep(1 * time.Second)
	fmt.Println(times)
}*/