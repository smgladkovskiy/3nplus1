package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

type sequence struct {
	initialNumber int64
	steps         int64
	currentNumber int64
	maxNumber     int64
}

type seqResult struct {
	number    int64
	steps     int64
	maxNumber int64
}

func main() {
	var (
		n    int64
		x    = int64(1)
		mn   = seqResult{}
		ms   = seqResult{}
		t    = time.Now()
		maxI = int64(10e6)
		res  = make(chan sequence)
		sem  = semaphore.NewWeighted(int64(10e3))
	)

	defer func() {
		close(res)
	}()

	go func() {
		for i := int64(1); i <= maxI; i++ {
			semGo(sem)

			go func(i int64) {
				s := sequence{initialNumber: i, steps: 0, currentNumber: i}
				s.iterate(res)
			}(i)
		}
	}()

checkCycle:
	for {
		select {
		case s := <-res:
			if ms.steps < s.steps {
				ms.maxNumber = s.maxNumber
				ms.number = s.initialNumber
				ms.steps = s.steps
			}

			if mn.maxNumber < s.maxNumber {
				mn.maxNumber = s.maxNumber
				mn.number = s.initialNumber
				mn.steps = s.steps
			}

			if n < s.initialNumber {
				n = s.initialNumber
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
		ms.number,
		ms.steps,
		mn.number,
		mn.maxNumber,
	)
}

func semGo(sem *semaphore.Weighted) {
	if sem.TryAcquire(1) {
		return
	}

	semGo(sem)
}

func (s *sequence) iterate(res chan sequence) {
	s.steps++

	if s.currentNumber%2 != 0 {
		s.currentNumber = s.currentNumber*3 + 1
	} else {
		s.currentNumber /= 2
	}

	if s.maxNumber < s.currentNumber {
		s.maxNumber = s.currentNumber
	}

	if s.currentNumber == 1 {
		res <- *s

		return
	}

	s.iterate(res)
}
