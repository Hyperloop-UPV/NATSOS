package generator

import (
	mrand "math/rand"
	"sync/atomic"
	"time"
)

var seedCounter atomic.Int64

// newRandomSeed returns a seed that won't be repited
func newRandomSeed() int64 {
	return time.Now().UnixNano() + seedCounter.Add(1)
}

// Generates a new random generator based in each seed
func newRNG() *mrand.Rand {
	return mrand.New(mrand.NewSource(newRandomSeed()))
}
