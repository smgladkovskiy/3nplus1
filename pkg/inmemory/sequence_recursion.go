package inmemory

import (
	"github.com/smgladkovskiy/3nplus1/pkg"
)

type RecursionSequence struct {
	numbersStorage pkg.NumbersStorageInterface
	resultsChan    chan pkg.SequenceResultInterface
}

func NewSequenceRecursion(iNumbers pkg.NumbersStorageInterface, res chan pkg.SequenceResultInterface) pkg.SequenceInterface {
	return &RecursionSequence{
		numbersStorage: iNumbers,
		resultsChan:    res,
	}
}

func (s *RecursionSequence) RunForNumber(i int64) {
	sr := &pkg.SequenceResult{
		InitialNumber: i,
		CurrentNumber: i,
		Numbers:       []int64{i},
	}
	s.resultsChan <- s.CollatzSequence(sr)
}

func (s *RecursionSequence) CollatzSequence(sr pkg.SequenceResultInterface) pkg.SequenceResultInterface {
	sr.CollatzIteration()

	if sr.IsEnd() || s.AlreadyCounted(sr) {
		return sr
	}

	return s.CollatzSequence(sr)
}

func (s *RecursionSequence) AlreadyCounted(sr pkg.SequenceResultInterface) bool {
	return s.numbersStorage.InMap(sr.GetCurrentNumber())
}
