package main

import (
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"jianshangquan.com/myapp/array"
	"jianshangquan.com/myapp/class"
	"jianshangquan.com/myapp/class/interfaces"
	"jianshangquan.com/myapp/conditional"
	"jianshangquan.com/myapp/filestream"
	"jianshangquan.com/myapp/lib"
	"jianshangquan.com/myapp/loop"
	"jianshangquan.com/myapp/maps"
	"jianshangquan.com/myapp/namespaces"
	"jianshangquan.com/myapp/passbyref"
	"jianshangquan.com/myapp/routines/basicroutines"
	routinechan "jianshangquan.com/myapp/routines/routine-chan"
	routinemutex "jianshangquan.com/myapp/routines/routine-mutex"
	"jianshangquan.com/myapp/slice"
	trycatch "jianshangquan.com/myapp/try-catch"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get caller info")
	}

	dir := filepath.Dir(filename)

	var waitGroup = &sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		var fileReader = filestream.NewFileStreamReader(path.Join(dir, "test.txt"))
		var fileWriter = filestream.NewFileStreamWriter(path.Join(dir, "test2.txt"))
		fileReader.PipeInChunkWithCallback(fileWriter, 1, filestream.WaitOnChunk(1))
		defer fileReader.Close()
		defer fileWriter.Close()
		waitGroup.Done()
	}()

	waitGroup.Wait()
	println("Main function starting...")

}

func tryCatchTest() {
	trycatch.TestTryCatch()
}

func passByRefTest() {
	passbyref.PassByRefTest()
}

func namespaceTest() {
	namespaces.Add(1, 2)
	namespaces.Cosine()
}

func routineMutex() {
	routinemutex.TestRoutineMutex()
}

func routineChan() {
	routinechan.TestRoutineChan()
	time.Sleep(10 * time.Second)
	println("Main function ending...")
}

func basicRoutine() {
	basicroutines.GoRoutine(5)

	time.Sleep(10 * time.Second)
	println("Main function ending...")
}

func basicSyntaxMain() {
	// This is a simple Go program that prints "Hello, World!" to the console.
	println("Hello, World!")

	println("\nArray")
	println((lib.AddData(1, 2)))
	println(array.List[0])

	println("\nSlice")
	println(slice.Slice[0])
	slice.ModifySlice()
	println(slice.Slice[0])

	println("\nConditional")
	println(conditional.IsPassed(60))

	println("\nMap")
	println(maps.Map["key"])

	println("\nLoop")
	println(loop.Loop())

	println("\nStruct")
	var car *class.Car = class.NewCar()
	println(car.Brand)
	println(car.Model)
	println(car.Year)
	println(car.Price)
	car.Run()

	println("\nInterface")
	var interfaceCar interfaces.Vehicle = class.NewCar()
	interfaceCar.Drive()
}
