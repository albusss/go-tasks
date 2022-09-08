package fastAndSlow

import "fmt"

func HappyNUmber() {
	num := 23
	sum := 0
	digit := 0
	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num = num / 10
	}

	fmt.Println(sum)
}
