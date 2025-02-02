package reading_time

import (
	"math"
	"time"
)

// Estimator holds the parameters to calculate the reading time.
type Estimator struct {
	WordsPerMinute    int
	SecondsPerLine    int
	BaseImageSeconds  int
	ImageSecondsDecay int
	ImageThreshold    int
}

// NewEstimator creates a new Estimator object with custom parameters.
func NewEstimator(wordsPerMinute, secondsPerLine, baseImageSeconds, imageSecondsDecay, imageThreshold int) *Estimator {
	return &Estimator{
		WordsPerMinute:    wordsPerMinute,
		SecondsPerLine:    secondsPerLine,
		BaseImageSeconds:  baseImageSeconds,
		ImageSecondsDecay: imageSecondsDecay,
		ImageThreshold:    imageThreshold,
	}
}

// StandardEstimator is a default estimator with common parameters for average reading speed.
var StandardEstimator = NewEstimator(200, 2, 12, 1, 10)

// FastEstimator for faster reading.
var FastEstimator = NewEstimator(250, 1, 10, 1, 8)

// SlowEstimator for slower reading.
var SlowEstimator = NewEstimator(150, 3, 15, 2, 12)

// CalculateReadingTime calculates the total reading time.
func (e *Estimator) CalculateReadingTime(totalWords, totalImages, totalLinesOfCode int) time.Duration {
	return e.CalculateWordsTime(totalWords) +
		e.CalculateImagesTime(totalImages) +
		e.CalculateCodeTime(totalLinesOfCode)
}

// CalculateWordsTime calculates reading time based on word count.
func (e *Estimator) CalculateWordsTime(totalWords int) time.Duration {
	if e.WordsPerMinute <= 0 {
		return 0
	}
	wordsPerSecond := 60.0 / float64(e.WordsPerMinute)
	return roundDurationToNearestSecond(float64(totalWords) * wordsPerSecond)
}

// CalculateImagesTime calculates additional time based on the number of images.
func (e *Estimator) CalculateImagesTime(totalImages int) time.Duration {
	var adjustmentTime float64
	for i := 0; i < totalImages; i++ {
		decayedTime := float64(e.BaseImageSeconds) - float64(i)*float64(e.ImageSecondsDecay)
		if i >= e.ImageThreshold {
			decayedTime = float64(e.BaseImageSeconds) - float64(e.ImageThreshold-1)*float64(e.ImageSecondsDecay)
		}
		if decayedTime > 0 {
			adjustmentTime += decayedTime
		}
	}
	return roundDurationToNearestSecond(adjustmentTime)
}

// CalculateCodeTime calculates additional time based on lines of code.
func (e *Estimator) CalculateCodeTime(totalLinesOfCode int) time.Duration {
	return roundDurationToNearestSecond(float64(totalLinesOfCode) * float64(e.SecondsPerLine))
}

// roundDurationToNearestSecond rounds the duration to the nearest second.
func roundDurationToNearestSecond(seconds float64) time.Duration {
	if seconds < 0 {
		return 0
	}
	return time.Duration(math.Round(seconds)) * time.Second
}
