package main

import (
	"fmt"
	"tasks/task"
	"time"
)

func main() {
	start := time.Now()
	// res := task.Descriminant(4, 4, 1)
	// res := task.Surface(5)
	res := task.RingS(3, 5)
	fmt.Println(res)
	timeElapsed := time.Since(start)
	fmt.Printf("it took %s\n", &timeElapsed)
}
