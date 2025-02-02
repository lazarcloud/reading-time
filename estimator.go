// package reading_time provides utilities to calculate reading time of a blog post
package reading_time

import (
	"time"
)

// Estimator holds the parameters to calculate the reading time
type Estimator struct {
	WordsPerMinute time.Duration
	SecondsPerLine time.Duration
	BaseImageTime  time.Duration
	ImageTimeDecay time.Duration
	ImageThreshold int
}

// NewEstimator creates a new Estimator object with custom parameters
func NewEstimator(wordsPerMinute, secondsPerLine, baseImageTime, imageTimeDecay time.Duration, imageThreshold int) *Estimator {
	return &Estimator{
		WordsPerMinute: wordsPerMinute,
		SecondsPerLine: secondsPerLine,
		BaseImageTime:  baseImageTime,
		ImageTimeDecay: imageTimeDecay,
		ImageThreshold: imageThreshold,
	}
}

// StandardEstimator is a default estimator with common parameters for average reading speed.
// It assumes 200 words per minute, 2 seconds per line of code, 12 seconds for each image,
// with a decay of 1 second per image until 10 images, after which the time per image becomes constant.
var StandardEstimator = NewEstimator(
	time.Duration(200),
	time.Duration(2),
	time.Duration(12),
	time.Duration(1),
	10,
)

// FastEstimator is designed for a faster reading experience.
// It assumes 250 words per minute, 1 second per line of code, 10 seconds base time for each image,
// with a decay of 1 second per image until 8 images, after which the time per image becomes constant.
var FastEstimator = NewEstimator(
	time.Duration(250),
	time.Duration(1),
	time.Duration(10),
	time.Duration(1),
	8,
)

// SlowEstimator is designed for a slower reading experience.
// It assumes 150 words per minute, 3 seconds per line of code, 15 seconds base time for each image,
// with a decay of 2 seconds per image until 12 images, after which the time per image becomes constant.
var SlowEstimator = NewEstimator(
	time.Duration(150),
	time.Duration(3),
	time.Duration(15),
	time.Duration(2),
	12,
)

// CalculateReadingTime calculates the total reading time and stores it in the Estimator object
func (e *Estimator) CalculateReadingTime(totalWords, totalImages, totalLinesOfCode int) time.Duration {
	readingTime := e.CalculateWordsTime(totalWords)
	readingTime += e.CalculateImagesTime(totalImages)
	readingTime += e.CalculateCodeTime(totalLinesOfCode)

	return readingTime
}

// CalculateWordsTime calculates the reading time based on the word count (words per minute)
func (e *Estimator) CalculateWordsTime(totalWords int) time.Duration {
	wordsReadingTime := time.Duration(totalWords) * time.Minute / e.WordsPerMinute
	return wordsReadingTime
}

// CalculateImagesTime calculates additional time based on the number of images in the post
func (e *Estimator) CalculateImagesTime(totalImages int) time.Duration {
	var adjustmentTime time.Duration
	for i := 1; i <= totalImages; i++ {
		if i <= e.ImageThreshold {
			adjustmentTime += e.BaseImageTime - time.Duration(i-1)*e.ImageTimeDecay
		} else {
			adjustmentTime += e.BaseImageTime - time.Duration(e.ImageThreshold-1)*e.ImageTimeDecay
		}
	}
	return adjustmentTime
}

// CalculateCodeTime calculates additional time based on the lines of code in the post
func (e *Estimator) CalculateCodeTime(totalLinesOfCode int) time.Duration {
	lineAdjustmentTime := time.Duration(totalLinesOfCode) * e.SecondsPerLine
	return lineAdjustmentTime
}
