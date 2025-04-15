package basicroutines

func routineJob(id int) {
	for {
		println("Routine job", id)
	}
}


func GoRoutine(routineCount int){
	for i := 0; i < routineCount; i++{
		go routineJob(i);
	}
}