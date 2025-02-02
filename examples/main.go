package main

import (
	"fmt"

	reading_time "github.com/lazarcloud/reading-time"
)

func main() {
	estimator := reading_time.StandardEstimator

	time := estimator.CalculateReadingTime(1000, 5, 20)

	fmt.Println(time)
	fmt.Println(time.Minutes())
	fmt.Println(time.Seconds())
}
