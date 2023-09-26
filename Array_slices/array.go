package main

func Sum(numbers []int) (sum_value int) {

	for _, j := range numbers {
		sum_value += j
	}
	return
}

func SumAll(myslices1, myslices2 []int) []int {
	x := Sum(myslices1)
	y := Sum(myslices2)

	return []int{x, y}

}

func SumTails(myslices1, myslices2 []int) []int {
	var sums []int
	ab := Sum(myslices1[1:])
	cd := Sum(myslices2[1:])

	sums = append(sums, ab, cd)

	return sums
}
