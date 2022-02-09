package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите количество скобок:")
	var count int
	fmt.Scan(&count)
	size := count * 2
	str := make([]string, size)
	add(count, count, str, 0)
}

func add(leftCount int, rightCount int, str []string, index int) {

	if leftCount < 0 || leftCount > rightCount{
		fmt.Println("Error:", str)
		return
	}
  
	if leftCount == 0 && rightCount == 0 {
		fmt.Println(str)
	} else {
		if leftCount > 0 {
			str[index] = "("
			add(leftCount-1, rightCount, str, index+1)
		}

		if leftCount < rightCount {
			str[index] = ")"
			add(leftCount, rightCount-1, str, index+1)
		}

	}

}
