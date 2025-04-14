package routinechan

import "time"



func routineJob(id int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- id*100 + i // send data
		time.Sleep(500 * time.Millisecond)
	}
}


func TestRoutineChan(){
	var ch = make(chan int)

	// Launch goroutines
	for i := 1; i <= 3; i++ {
		go routineJob(i, ch)
	}

	// Receive from channel
	for i := 0; i < 15; i++ {
		val := <-ch
		println("Received:", val)
	}

	println("All data received!")
}