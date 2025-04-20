package passfunction

func Run[T any](callback *func(value T) T) T {
	// Call the callback function
	var name T
	return (*callback)(name)
}

func TestRun() {
	// Define a callback function
	callback := func(value string) string {
		return "Hello, " + value
	}

	// Call the Run function with the callback
	result := Run(&callback)

	// Print the result
	println(result)
}
