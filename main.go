package main

import "fmt"

const (
	limit = 823
)

func getNextDay(year, day int) int {
	step := 1
	if year%4 == 0 {
		step = 2
	}

	result := day + step
	if result > 7 {
		result = result - 7
	}

	return result
}

func main() {
	weekDay := 1
	year := 2021
	endYear := year + limit

	result := []int{}

	for i := year; i < endYear; i++ {
		if i > year {
			weekDay = getNextDay(i, weekDay)
		}

		if i%4 != 0 && weekDay == 1 {
			result = append(result, i)
		}
	}

	fmt.Printf("Hay %d años entre el %d y el %d que tienen lunes 1 de febrero con 28 días:\n%v\n",
		len(result), year, endYear, result)
}
