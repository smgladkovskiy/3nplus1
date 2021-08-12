package main

import (
	"context"
	"flag"
	"log"
	"math"
	"runtime"
	"time"

	"github.com/marusama/semaphore/v2"
	"github.com/smgladkovskiy/3nplus1/models"
)

func main() {
	flag.Parse()
	cpuProfile()

	var (
		n               int64
		x               = int64(1)
		mn              = models.SeqResult{}
		t               = time.Now()
		maxI            = int64(math.Pow10(*maxIPower))
		res             = make(chan models.Sequence)
		semConcurrent   = runtime.NumCPU() * 4
		sem             = semaphore.New(semConcurrent)
		iteratedNumbers = models.NewAllIteratedNumbers(maxI)
	)

	runtime.GOMAXPROCS(semConcurrent)

	defer close(res)

	log.Printf("Lets iterate numbers from 1 to 10e%d in %d parallel processes\n", *maxIPower, semConcurrent)

	ctx := context.Background()

	go func() {
		for i := int64(1); i <= maxI; i++ {
			_ = sem.Acquire(ctx, 1)

			go func() {
				s := models.Sequence{InitialNumber: i, Steps: 0, CurrentNumber: i, Numbers: []int64{i}}
				s.Iterate(&iteratedNumbers, res)
			}()
		}
	}()

	for s := range res {
		go iteratedNumbers.AddNumbers(s.Numbers)

		if mn.MaxNumber < s.MaxNumber {
			mn.MaxNumber = s.MaxNumber
			mn.Number = s.InitialNumber
			mn.Steps = s.Steps
		}

		if n < s.InitialNumber {
			n = s.InitialNumber
		}

		if x == maxI {
			break
		}

		x++

		sem.Release(1)
	}

	endTime := time.Since(t).String()

	memProfile()

	log.Println("Iterations are done. Preparing results...")

	log.Printf(
		"Results:\nmax number: %d\ntime: %s\nnumbers in all iterations %d\nmax reached value for number %d is %d\n",
		n,
		endTime,
		iteratedNumbers.Numbers.Size(),
		mn.Number,
		mn.MaxNumber,
	)
}
