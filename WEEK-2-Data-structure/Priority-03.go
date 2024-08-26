package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", sum(data1))       //output 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     //output 5.46
	fmt.Printf("median: %.2f\n", median(data1)) //output 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     //output 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("sum: %.2f\n", sum(data2))       //output 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     //output 7.92
	fmt.Printf("median: %.2f\n", median(data2)) //output 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     //output 1.00
}

func sum(data []float64) float64 {
	total := 0.0
	for _, value := range data {
		total += value
	}
	return total
}

func mean(data []float64) float64 {
	total := sum(data)
	return total / float64(len(data))
}

func median(data []float64) float64 {
	n := len(data)
	sort.Float64s(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	} else {
		return data[n/2]
	}
}

func mode(data []float64) float64 {
	frequency := make(map[float64]int)
	maxFreq := 0
	mode := math.NaN()

	for _, value := range data {
		frequency[value]++
		if frequency[value] > maxFreq {
			maxFreq = frequency[value]
			mode = value
		}
	}


	for key, value := range frequency {
		if value == maxFreq && (math.IsNaN(mode) || key < mode) {
			mode = key
		}
	}

	return mode
}
