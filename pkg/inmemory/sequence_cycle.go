package inmemory

import (
	"github.com/smgladkovskiy/3nplus1/pkg"
)

type CycleSequence struct {
	numbersStorage pkg.NumbersStorageInterface
	resultsChan    chan pkg.SequenceResultInterface
}

func NewSequenceWithCycle(iNumbers pkg.NumbersStorageInterface, res chan pkg.SequenceResultInterface) pkg.SequenceInterface {
	return &CycleSequence{
		numbersStorage: iNumbers,
		resultsChan:    res,
	}
}

func (s *CycleSequence) RunForNumber(i int64) {
	sr := &pkg.SequenceResult{
		InitialNumber: i,
		CurrentNumber: i,
		Numbers:       []int64{i},
	}
	s.resultsChan <- s.CollatzSequence(sr)
}

func (s *CycleSequence) CollatzSequence(sr pkg.SequenceResultInterface) pkg.SequenceResultInterface {
	for {
		sr.CollatzIteration()

		if sr.IsEnd() || s.AlreadyCounted(sr) {
			return sr
		}
	}
}

func (s *CycleSequence) AlreadyCounted(sr pkg.SequenceResultInterface) bool {
	return s.numbersStorage.InMap(sr.GetCurrentNumber())
}
