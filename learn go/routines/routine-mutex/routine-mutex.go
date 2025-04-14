package routinemutex

import (
	"sync"
	"time"
)


var counter int = 0;

func routineJob(id int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		mu.Lock()   // lock before modifying
		counter++
		println("Routine", id, "incremented counter to", counter)
		mu.Unlock() // unlock after modification

		time.Sleep(300 * time.Millisecond)
	}
}


func TestRoutineMutex(){
	
	var mutex = sync.Mutex{}
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go routineJob(i, &mutex, &wg)
	}

	wg.Wait() // wait for all to finish
	println("Final counter:", counter)
}