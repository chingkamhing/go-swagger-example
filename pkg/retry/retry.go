package retry

import (
	"math/rand"
	"time"
)

// Note: this helper function is copied from https://gist.github.com/sascha-andres/d1f11fb9bc6abc4f07b4118839b29d7f

// jitterDivisor control how large of the random jitter time
const jitterDivisor = 4

// Stop is a retry stopper error; i.e. if the retry function want to stop retrying before attempts, return this Sop error instead
type Stop struct {
	Err error
}

// Error implement error interface
func (s Stop) Error() string {
	return s.Err.Error()
}

// Config is a retry config that hold retry config settings
type Config struct {
	attempts int
	sleep    time.Duration
}

func init() {
	// random seed
	rand.Seed(time.Now().UnixNano())
}

// Do retry calling retryFunc for attempts number of times if it return error. The retryFunc may quit retrying by returning Stop error instead.
func Do(attempts int, sleep time.Duration, retryFunc func(int) error) (err error) {
	currSleep := sleep
	for i := 0; i < attempts; i++ {
		err = retryFunc(i)
		if err != nil {
			// Q: Stop error? return the original error for later checking if it is Stop error
			if s, ok := err.(Stop); ok {
				return s.Err
			}
			// add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(currSleep)) - int64(currSleep)/2)
			time.Sleep(currSleep + jitter/jitterDivisor)
			currSleep = 2 * currSleep
		} else {
			return nil
		}
	}
	return err
}
