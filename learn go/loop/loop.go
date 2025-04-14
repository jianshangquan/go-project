package loop

import "jianshangquan.com/myapp/slice"



func Loop() int {

	var loopedCount int = 0;

	for i := 0; i < 10; i++ {
		println(i)
		loopedCount++;
	}

	// Loop through the array using range
	for index, value := range slice.Slice {
		println("Index: %d, Value: %d\n", index, value)
		loopedCount++;
	}
	return loopedCount;
}