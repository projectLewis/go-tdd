package main

// Add adds multiple numbers
func Add(num ...int) (sum int) {
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return
}

func main() {
	Add(2, 2)
}
