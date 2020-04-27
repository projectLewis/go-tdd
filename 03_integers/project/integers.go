package main

// Add adds multiple numbers
func Add(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func main() {
	Add(2, 2)
}
