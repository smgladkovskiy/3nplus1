package cmd

import (
	"context"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/marusama/semaphore/v2"
	"github.com/smgladkovskiy/3nplus1/pkg"
	"github.com/smgladkovskiy/3nplus1/pkg/inmemory"
	"github.com/spf13/cobra"
)

func runIterator(_ *cobra.Command, _ []string) {
	cpuProfile()

	var (
		n               int64
		x               = int64(1)
		mn              = newSequenceResult()
		t               = time.Now()
		maxI            = int64(math.Pow10(maxIPower))
		res             = make(chan pkg.SequenceResultInterface)
		semConcurrent   = runtime.NumCPU() * 4
		sem             = semaphore.New(semConcurrent)
		iteratedNumbers pkg.NumbersStorageInterface
	)

	iteratedNumbers = inmemory.NewAllIteratedNumbers(maxI)

	// runtime.GOMAXPROCS(semConcurrent)

	defer close(res)

	log.Printf("Lets iterate numbers from 1 to 10e%d in %d parallel processes\n", maxIPower, semConcurrent)

	ctx := context.Background()

	go func() {
		seq := inmemory.NewSequence(iteratedNumbers, res)

		for i := int64(1); i <= maxI; i++ {
			_ = sem.Acquire(ctx, 1)

			go seq.RunForNumber(i)
		}
	}()

	for s := range res {
		go iteratedNumbers.AddNumbers(s.GetNumbers())

		if mn.GetMaxNumber() < s.GetMaxNumber() {
			mn = s
		}

		if n < s.GetInitialNumber() {
			n = s.GetInitialNumber()
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
		iteratedNumbers.Amount(),
		mn.GetInitialNumber(),
		mn.GetMaxNumber(),
	)
}

func newSequenceResult() pkg.SequenceResultInterface {
	return &pkg.SequenceResult{}
}

func cpuProfile() {
	if cpuprofile {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}
}

func memProfile() {
	if cpuprofile {
		pprof.StopCPUProfile()
	}

	if memprofile {
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}

		runtime.GC() // get up-to-date statistics

		if err := pprof.WriteHeapProfile(f); err != nil {
			_ = f.Close()

			log.Fatal("could not write memory profile: ", err)
		}

		_ = f.Close()
	}
}
