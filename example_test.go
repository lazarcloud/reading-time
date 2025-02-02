package reading_time_test

import (
	"fmt"

	reading_time "github.com/lazarcloud/reading-time"
)

func ExampleEstimator_CalculateWordsTime() {
	// Custom estimator with 250 words per minute
	customEstimator := reading_time.NewEstimator(250, 2, 12, 1, 10)
	fmt.Println(customEstimator.CalculateWordsTime(100))
	// Output: 24s
}

func ExampleEstimator_CalculateWordsTime_negative() {
	// Custom estimator with negative words per minute
	customEstimator := reading_time.NewEstimator(-250, 2, 12, 1, 10)
	fmt.Println(customEstimator.CalculateWordsTime(100))
	// Output: 0s
}

func ExampleEstimator_CalculateImagesTime_basic() {
	// Custom estimator with 10 seconds base image time and decay of 2 seconds per image
	customEstimator := reading_time.NewEstimator(200, 2, 10, 2, 5)
	fmt.Println(customEstimator.CalculateImagesTime(5))
	// Output: 30s
}

func ExampleEstimator_CalculateImagesTime_more() {
	// Custom estimator with 10 seconds base image time and decay of 2 seconds per image
	customEstimator := reading_time.NewEstimator(200, 2, 10, 2, 5)

	// Test case with a larger number of images
	fmt.Println(customEstimator.CalculateImagesTime(10))
	// Output: 40s
}

func ExampleEstimator_CalculateCodeTime() {
	// Custom estimator with 3 seconds per line of code
	customEstimator := reading_time.NewEstimator(200, 3, 12, 1, 10)
	fmt.Println(customEstimator.CalculateCodeTime(10))
	// Output: 30s
}

func ExampleEstimator_CalculateReadingTime_basic() {
	// Custom estimator with 200 words per minute, 2 seconds per line, and base image time of 15 seconds
	customEstimator := reading_time.NewEstimator(200, 2, 15, 1, 8)
	readingTime := customEstimator.CalculateReadingTime(100, 5, 10)
	fmt.Println(readingTime)
	// Output: 1m55s
}

func ExampleEstimator_CalculateReadingTime_faster() {
	// Example with a faster reading custom estimator
	fasterEstimator := reading_time.NewEstimator(300, 1, 10, 1, 6)
	readingTimeFast := fasterEstimator.CalculateReadingTime(100, 5, 10)
	fmt.Println(readingTimeFast)
	// Output: 1m10s
}
