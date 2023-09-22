package main

func Sum(numbers [5]int) (sum_value int) {

	for _, j := range numbers {
		sum_value += j
	}
	return
}
