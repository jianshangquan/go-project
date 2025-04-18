package trycatch


func TestTryCatch(){
	defer func() {
		println("defer 2")
		var recoverResult = recover()
		println("recover", recoverResult)
	}()
	println("func")
	panic("error")
}