package reading_time

import (
	"fmt"
	"time"
)

func ExampleCalculateWordsTime() {
	// Custom estimator with 250 words per minute
	customEstimator := NewEstimator(
		time.Duration(250),
		time.Duration(2),
		time.Duration(12),
		time.Duration(1),
		10,
	)
	fmt.Println(customEstimator.CalculateWordsTime(100))
	// Output: 24m0s (since 100 words at 250 words per minute equals 24 minutes)
}

func ExampleCalculateImagesTime() {
	// Custom estimator with 10 seconds base image time and decay of 2 seconds per image
	customEstimator := NewEstimator(
		time.Duration(200),
		time.Duration(2),
		time.Duration(10),
		time.Duration(2),
		5,
	)
	fmt.Println(customEstimator.CalculateImagesTime(5))
	// Output: 40s (5 images, each with 10 seconds base time and decay of 2 seconds per image)

	// Test case with a larger number of images
	fmt.Println(customEstimator.CalculateImagesTime(10))
	// Output: 90s (10 images with decay for the first 5 images and constant time for the remaining 5)
}

func ExampleCalculateCodeTime() {
	// Custom estimator with 3 seconds per line of code
	customEstimator := NewEstimator(
		time.Duration(200),
		time.Duration(3),
		time.Duration(12),
		time.Duration(1),
		10,
	)
	fmt.Println(customEstimator.CalculateCodeTime(10))
	// Output: 30s (10 lines of code, each with 3 seconds per line)
}

func ExampleCalculateReadingTime() {
	// Custom estimator with 200 words per minute, 2 seconds per line, and base image time of 15 seconds
	customEstimator := NewEstimator(
		time.Duration(200),
		time.Duration(2),
		time.Duration(15),
		time.Duration(1),
		8,
	)
	readingTime := customEstimator.CalculateReadingTime(100, 5, 10)
	fmt.Println(readingTime)
	// Output: 31m40s (total of word time, image time, and code time with custom estimator)

	// Example with a faster reading custom estimator
	fasterEstimator := NewEstimator(
		time.Duration(300),
		time.Duration(1),
		time.Duration(10),
		time.Duration(1),
		6,
	)
	readingTimeFast := fasterEstimator.CalculateReadingTime(100, 5, 10)
	fmt.Println(readingTimeFast)
	// Output: 29m0s (faster reading speed with custom values)
}
