package main

import "fmt"

func getvolume(input []int) int {
	sumtotal := 0

	high := 0
	highvalue := input[high]

	for i := 1; i < len(input); i++ {

		if i == len(input) {
			highreverse := len(input)
			highvaluereverse := input[highreverse]
			if(highvalue < highvaluereverse) {
				highvaluereverse = highvalue
			}

			for j := len(input) - 1; j > high; j++ {
				sumtotal -= highvalue - input[j]

				current := input[j]
				if current < highvaluereverse {
					sumtotal += highvaluereverse - current
				} else if current > highvaluereverse {
					highreverse = i
					highvaluereverse = input[i]
				}
			}

			break
		}

		current := input[i]
		if current < highvalue {
			sumtotal += highvalue - current
		} else if current > highvalue {
			high = i
			highvalue = input[i]
		}
	}
	return sumtotal
}

func main() {
	// fmt.Println("4, 2, 3, 2, 5, 3, 2, 4, 2, 5: ",getvolume([]int{4, 2, 3, 2, 5, 3, 2, 4, 2, 5}))
	// fmt.Println("1, 2, 3, 4, 5: ", getvolume([]int{1, 2, 3, 4, 5}))
	fmt.Println("5, 4, 3, 2, 1: ", getvolume([]int{5, 4, 3, 2, 1}))
	fmt.Println("1, 2, 3, 2, 1: ",getvolume([]int{1, 2, 3, 2, 1}))
	// fmt.Println("5, 0, 5: ",getvolume([]int{5, 0, 5}))
}
