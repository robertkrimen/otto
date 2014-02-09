package otto

import (
	"fmt"
	"testing"
	"time"
	. "./terst"
)

func limitedLoop(otto *Otto, d time.Duration) (iterations int64) {
	toRun := fmt.Sprintf(`
	  var d = new Date();
		var iters = 0;
		var deadline = d.getTime() + %d;
		while (new Date().getTime() < deadline) {
			iters++;
		}
		iters;
	`, d/time.Millisecond)

	result, err := otto.Run(toRun)

	Is(err, nil)

	iterations, err = result.ToInteger()

	Is(err, nil)

	return
}

func TestLimitedBurstEvaluation(t *testing.T) {
	Terst(t)

	otto := New()

	otto.Limits.Evaluation.Burst = 10

	otto.Limits.Evaluation.PerSecond = 1000

	firstResult := limitedLoop(otto, time.Millisecond*100)

	otto.Limits.Evaluation.PerSecond = 10000

	secondResult := limitedLoop(otto, time.Millisecond*100)

	// Check that the second result is about 10 times bigger than the first, since it can evaluate 10 times faster and the burst is quite small.
	Compare(secondResult, ">", firstResult*8)
	Compare(secondResult, "<", firstResult*12)
}

func TestBurstEvaluation(t *testing.T) {
	Terst(t)

	otto := New()

	otto.Limits.Evaluation.Burst = 10000000

	otto.Limits.Evaluation.PerSecond = 1000

	firstResult := limitedLoop(otto, time.Millisecond*100)

	otto.Limits.Evaluation.PerSecond = 10000

	secondResult := limitedLoop(otto, time.Millisecond*100)

	// Check that the two results are about the same, since both should have fit within the generous burst we set.
	Compare(secondResult*10, ">", firstResult*8)
	Compare(secondResult*10, "<", firstResult*12)
}

func TestEvaluationsPerSecond(t *testing.T) {
	Terst(t)

	otto := New()

	otto.Limits.Evaluation.PerSecond = 1000

	firstResult := limitedLoop(otto, time.Millisecond*100)

	otto.Limits.Evaluation.PerSecond = 10000

	secondResult := limitedLoop(otto, time.Millisecond*100)

	// Check that the second result is about 10 times bigger than the first, since it can evaluate 10 times faster.
	Compare(secondResult, ">", firstResult*8)
	Compare(secondResult, "<", firstResult*12)
}
