package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}

		switch i {
		case 0:
			fmt.Println("Zero")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		}
	}

}
