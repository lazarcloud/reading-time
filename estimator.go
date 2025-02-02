package reading_time

import (
	"time"
)

// Estimator holds the parameters to calculate the reading time
type Estimator struct {
	WordsPerMinute int
	SecondsPerLine time.Duration
	BaseImageTime  time.Duration
	ImageTimeDecay time.Duration
	ImageThreshold int
}

// NewEstimator creates a new Estimator object with custom parameters
func NewEstimator(wordsPerMinute int, secondsPerLine, baseImageTime, imageTimeDecay time.Duration, imageThreshold int) *Estimator {
	return &Estimator{
		WordsPerMinute: wordsPerMinute,
		SecondsPerLine: secondsPerLine,
		BaseImageTime:  baseImageTime,
		ImageTimeDecay: imageTimeDecay,
		ImageThreshold: imageThreshold,
	}
}

// StandardEstimator is a default estimator with common parameters for average reading speed.
var StandardEstimator = NewEstimator(200, time.Duration(2), time.Duration(12), time.Duration(1), 10)

// FastEstimator for faster reading
var FastEstimator = NewEstimator(250, time.Duration(1), time.Duration(10), time.Duration(1), 8)

// SlowEstimator for slower reading
var SlowEstimator = NewEstimator(150, time.Duration(3), time.Duration(15), time.Duration(2), 12)

// CalculateReadingTime calculates the total reading time
func (e *Estimator) CalculateReadingTime(totalWords, totalImages, totalLinesOfCode int) time.Duration {
	readingTime := e.CalculateWordsTime(totalWords)
	readingTime += e.CalculateImagesTime(totalImages)
	readingTime += e.CalculateCodeTime(totalLinesOfCode)

	return readingTime
}

// CalculateWordsTime calculates reading time based on word count
func (e *Estimator) CalculateWordsTime(totalWords int) time.Duration {
	wordsPerSecond := time.Minute / time.Duration(e.WordsPerMinute)
	wordsReadingTime := time.Duration(totalWords) * wordsPerSecond
	return wordsReadingTime
}

// CalculateImagesTime calculates additional time based on the number of images
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

// CalculateCodeTime calculates additional time based on lines of code
func (e *Estimator) CalculateCodeTime(totalLinesOfCode int) time.Duration {
	lineAdjustmentTime := time.Duration(totalLinesOfCode) * e.SecondsPerLine
	return lineAdjustmentTime
}
