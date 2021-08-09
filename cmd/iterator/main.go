package main

import (
	"flag"
	"fmt"
	"math"
	_ "net/http/pprof"
	"time"

	"github.com/smgladkovskiy/3nplus1/pkg/models"
	"golang.org/x/sync/semaphore"
)

func main() {
	flag.Parse()
	cpuProfile()

	var (
		n             int64
		x             = int64(1)
		mn            = models.SeqResult{}
		ms            = models.SeqResult{}
		t             = time.Now()
		startI        = int64(math.Pow(10, *minIPower))
		maxI          = int64(math.Pow(10, *maxIPower))
		res           = make(chan models.Sequence)
		semConcurrent = int64(10e5)
		sem           = semaphore.NewWeighted(semConcurrent)
		passedNumbers = models.NewCheckedNumbers()
	)

	defer close(res)

	fmt.Printf("Lets iterate numbers from %d to %d in %d parallel processes", startI, maxI, semConcurrent)

	go func() {
		for i := startI; i <= maxI; i++ {
			for {
				if sem.TryAcquire(1) {
					break
				}
			}

			s := models.Sequence{InitialNumber: i, Steps: 0, CurrentNumber: i}
			go s.Iterate(&passedNumbers, res)
		}
	}()

checkCycle:
	for {
		select {
		case s := <-res:
			passedNumbers.Set(s.InitialNumber)
			if ms.Steps < s.Steps {
				ms.MaxNumber = s.MaxNumber
				ms.Number = s.InitialNumber
				ms.Steps = s.Steps
			}

			if mn.MaxNumber < s.MaxNumber {
				mn.MaxNumber = s.MaxNumber
				mn.Number = s.InitialNumber
				mn.Steps = s.Steps
			}

			if n < s.InitialNumber {
				n = s.InitialNumber
			}

			if x == maxI {
				break checkCycle
			}

			x++
			sem.Release(1)
		default:
			_ = struct{}{}
		}
	}

	fmt.Printf(
		"max number: %d\ntime: %s\nmax steps for number %d is %d\nmax reached value for number %d is %d\n",
		n,
		time.Since(t).String(),
		ms.Number,
		ms.Steps,
		mn.Number,
		mn.MaxNumber,
	)

	memProfile()
}
