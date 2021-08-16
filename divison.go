package main

import "fmt"

func main() {
	divisible()
}

func divisible() {
	 count := 0
	 var numberList [] int

	 for i := 1; i < 100; i++ {
		 if i%3 == 0 && i%5 == 0 {
			 numberList = append(numberList, i)
			 count++
		 }
	 }

	fmt.Println(count)
	
	fmt.Println(numberList)


}