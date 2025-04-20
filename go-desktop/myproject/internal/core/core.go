package core

type CoreAPI struct{}

func (c *CoreAPI) Add(a, b int) int {
	return Add(a, b)
}

func Add(a, b int) int {
	return a + b
}
