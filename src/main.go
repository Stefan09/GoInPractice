package main

import "fmt"

func main() {
	for i := 50; i <= 99; i++ {
		if (i%3 == 0 || i%10 == 3) && (i%7 == 0 || i%10 == 7 || i/10 == 7) {
			fmt.Printf("%s ", "QQNews")
		} else if i%3 == 0 || i%10 == 3 {
			fmt.Printf("%s ", "QQ")
		} else if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Printf("%s ", "News")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
