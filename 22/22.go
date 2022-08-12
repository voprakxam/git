package main

import "fmt"

var msg string

func init() {
	msg = "from init()"
}

func main() {
	message := "123214"

	fmt.Println(message)
	printMessage(&message)

	fmt.Println(msg)
	fmt.Println(findMin(1, 2, 3, 4, -5, 6, 7))
	func() {
		fmt.Println("анонимная функция")
	}()
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())

}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func increment2() int {
	count := 0
	count++
	return count
}
func printMessage(message *string) {
	*message += "sdfsdfsdf"
	fmt.Println(message)
}
