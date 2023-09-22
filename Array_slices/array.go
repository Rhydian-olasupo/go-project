package main

func Sum(numbers [5]int) (sum_value int) {

	for i := 0; i < 5; i++ {
		sum_value += numbers[i]
	}
	return
}
